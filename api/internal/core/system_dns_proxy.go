package core

import (
	"fmt"
	"gen_gin_tpl/pkg/logger/log"
	"github.com/miekg/dns"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type DNSProxy struct {
	originalDNS []string
	hosts       map[string]string
}

// handleDNS 处理 DNS 请求
func (r *DNSProxy) handleDNS(w dns.ResponseWriter, msg *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(msg)
	upstream := "114.114.114.114:53"

	for _, q := range msg.Question {
		hostKey := strings.TrimSuffix(strings.ToLower(q.Name), ".")
		switch q.Qtype {
		case dns.TypeA:
			if ip, ok := r.hosts[hostKey]; ok {
				rr, _ := dns.NewRR(q.Name + " A " + ip)
				m.Answer = append(m.Answer, rr)
				log.Debug().Msgf("[命中本地DNS映射] %s -> %s", hostKey, ip)
			} else {
				// 转发到上游
				resp, err := dns.Exchange(msg, upstream)
				if err == nil {
					m.Answer = append(m.Answer, resp.Answer...)
				} else {
					m.Rcode = dns.RcodeNameError
					log.Debug().Err(err).Msgf("[上游DNS查询失败] %s", hostKey)
				}
			}
		default:
			// 其它类型请求直接转发
			resp, err := dns.Exchange(msg, upstream)
			if err == nil {
				m.Answer = append(m.Answer, resp.Answer...)
			} else {
				m.Rcode = dns.RcodeNameError
				log.Debug().Err(err).Msgf("[上游DNS查询失败-非A记录] %s", hostKey)
			}
		}
	}
	_ = w.WriteMsg(m)
}

// StartDnsServer 启动 DNS 服务器
func (r *DNSProxy) StartDnsServer(domain, ip string) {
	if r.hosts == nil {
		r.hosts = make(map[string]string)
	}
	// 注意：DNS 查询时域名会带 . 结尾，所以存的时候不带点，查的时候统一处理

	r.hosts[strings.ToLower(domain)] = ip

	// 启动 DNS 服务
	udpServer := &dns.Server{Addr: ":53", Net: "udp"}
	tcpServer := &dns.Server{Addr: ":53", Net: "tcp"}

	dns.HandleFunc(".", r.handleDNS)

	go func() {
		log.Info().Msg("[DNS代理] 启动本地 DNS 服务 (UDP) on *:53")
		if err := udpServer.ListenAndServe(); err != nil {
			log.Error().Err(err).Msg("本地 UDP DNS 启动失败！")
		}
	}()
	go func() {
		log.Info().Msg("[DNS代理] 启动本地 DNS 服务 (TCP) on *:53")
		if err := tcpServer.ListenAndServe(); err != nil {
			log.Error().Err(err).Msg("本地 TCP DNS 启动失败！")
		}
	}()
	time.Sleep(500 * time.Millisecond)
	log.Info().Msgf("[DNS代理] 已注册临时域名映射: %s -> %s", domain, ip)
}

// SetLocalDNS 设置系统 DNS 为 127.0.0.1
func (r *DNSProxy) SetLocalDNS() error {
	switch runtime.GOOS {
	case "windows":
		ifaces, _ := net.Interfaces()
		var target string
		for _, iface := range ifaces {
			if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 {
				target = iface.Name
				break
			}
		}
		if target == "" {
			return fmt.Errorf("没有找到可用网卡")
		}
		cmd := exec.Command("netsh", "interface", "ipv4", "set", "dns", `name="`+target+`"`, "static", "127.0.0.1")
		log.Debug().Str("Name", target).Str("OS", runtime.GOOS).Str("cmd", strings.Join(cmd.Args, " ")).Msg("系统代理DNS设置")
		if err := cmd.Run(); err != nil {
			return err
		}
	case "darwin":
		// macOS 可以列出所有网络服务再选第一个启用的
		cmd := exec.Command("networksetup", "-listallnetworkservices")
		out, _ := cmd.Output()
		services := strings.Split(string(out), "\n")
		var target string
		for _, s := range services {
			s = strings.TrimSpace(s)
			if s == "" || strings.HasPrefix(s, "*") { // 忽略注释或空行
				continue
			}
			target = s
			break
		}
		if target == "" {
			return fmt.Errorf("没有找到网络服务")
		}
		cmd = exec.Command("networksetup", "-setdnsservers", target, "127.0.0.1")
		log.Debug().Str("Name", target).Str("OS", runtime.GOOS).Str("cmd", strings.Join(cmd.Args, " ")).Msg("系统代理DNS设置")
		if err := cmd.Run(); err != nil {
			return err
		}
	case "linux":
		backup, _ := os.ReadFile("/etc/resolv.conf")
		r.originalDNS = strings.Split(string(backup), "\n")
		_ = os.WriteFile("/etc/resolv.conf", []byte("nameserver 127.0.0.1\n"), 0644)
	default:
		return fmt.Errorf("不支持的操作系统")
	}
	log.Info().Msg("[DNS代理] 系统 DNS 设置为 127.0.0.1")
	return nil
}

// RestoreDNS 恢复系统 DNS
//
// 说明:
//   - 恢复系统 DNS 为之前的配置。
//     c := make(chan os.Signal, 1)
//     signal.Notify(c, os.Interrupt, syscall.SIGTERM)
//     <-c
//     RestoreDNS()
func (r *DNSProxy) RestoreDNS() {
	switch runtime.GOOS {
	case "windows":
		ifaces, _ := net.Interfaces()
		var target string
		for _, iface := range ifaces {
			if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 {
				target = iface.Name
				break
			}
		}
		if target == "" {
			log.Info().Msg("没有找到网络服务，跳过恢复")
			return
		}
		cmd := exec.Command("netsh", "interface", "ipv4", "set", "dns", target, "dhcp")
		log.Debug().Str("Name", target).Str("OS", runtime.GOOS).Str("cmd", strings.Join(cmd.Args, " ")).Msg("[DNS代理] 系统代理DNS恢复")
		if err := cmd.Run(); err != nil {
			log.Error().Err(err).Msg("系统代理DNS设置失败")
		}
	case "darwin":
		cmd := exec.Command("networksetup", "-listallnetworkservices")
		out, _ := cmd.Output()
		services := strings.Split(string(out), "\n")
		var target string
		for _, s := range services {
			s = strings.TrimSpace(s)
			if s == "" || strings.HasPrefix(s, "*") {
				continue
			}
			target = s
			break
		}
		if target == "" {
			log.Info().Msg("没有找到网络服务，跳过恢复")
			return
		}
		cmd = exec.Command("networksetup", "-setdnsservers", target, "Empty")
		log.Debug().Str("Name", target).Str("OS", runtime.GOOS).Str("cmd", strings.Join(cmd.Args, " ")).Msg("[DNS代理] 系统代理DNS恢复")
		if err := cmd.Run(); err != nil {
			log.Error().Err(err).Msg("系统代理DNS设置失败")
		}
	case "linux":
		if len(r.originalDNS) > 0 {
			_ = os.WriteFile("/etc/resolv.conf", []byte(strings.Join(r.originalDNS, "\n")), 0644)
		}
	default:
		log.Error().Msgf("不支持的操作系统：%s", runtime.GOOS)
		return
	}
	log.Info().Msg("[DNS代理] 恢复系统 DNS 代理")
}
