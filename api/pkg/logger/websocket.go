package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gen_gin_tpl/pkg/variable"
	"github.com/gorilla/websocket"
	"sync"
)

var logQueue = make(chan []byte, 1000)

// logEvent 定义日志结构体
type logEvent struct {
	Time    string `json:"time"`
	Level   string `json:"level"`
	Caller  string `json:"caller"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func FormatLogMessage(p []byte) string {
	if CurrentCfg.FormatJSON {
		return string(p)
	}

	var event logEvent
	dec := json.NewDecoder(bytes.NewReader(p))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&event); err != nil {
		return string(p)
	}

	s := fmt.Sprintf("%s [%s] > %s", event.Time, event.Level, event.Message)
	if event.Error != "" {
		s += fmt.Sprintf(" error=\"%s\"", event.Error)
	}
	return s
}

// 客户端结构，带写队列和关闭通道
type client struct {
	conn   *websocket.Conn
	sendCh chan []byte
}

type hub struct {
	clients map[*client]bool
	lock    sync.RWMutex
}

var LogHub = &hub{
	clients: make(map[*client]bool),
}

// SendMessage 发送消息给所有客户端
// 每个 client 都有自己的 writer goroutine，负责写入消息
// 这里我们直接将消息发送到每个 client 的 sendCh 通道
// 每个 client 独立 writer goroutine 负责从 sendCh 读取消息并写入 WebSocket 连接
// 如果队列满了，我们可以选择踢掉 client，避免阻塞
// 注意：这里的 sendCh 是无缓冲的，所以如果 client 没有及时读取，会导致阻塞
// 如果需要缓冲，可以考虑使用带缓冲的 channel
// 参数：
//   - msg string: 要发送的消息
//
// 返回值：
//   - 无
func (h *hub) SendMessage(msg string) {
	h.broadcast([]byte(msg))
}

// AddClient 添加新的客户端
// 每个 client 都有自己的 writer goroutine，负责写入消息
// 参数：
//   - conn *websocket.Conn: WebSocket 连接
//
// 返回值：
//   - *client: 新添加的客户端
func (h *hub) AddClient(conn *websocket.Conn) *client {
	c := &client{
		conn:   conn,
		sendCh: make(chan []byte, 256),
	}

	h.lock.Lock()
	h.clients[c] = true
	h.lock.Unlock()

	// 每个 client 独立 writer goroutine
	go func() {
		defer func() {
			conn.Close()
			h.RemoveClient(c)
		}()
		for msg := range c.sendCh {
			if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				break
			}
		}
	}()

	return c
}

// RemoveClient 移除客户端
// 参数：
//   - c *client: 要移除的客户端
//
// 返回值：
//   - 无
func (h *hub) RemoveClient(c *client) {
	h.lock.Lock()
	defer h.lock.Unlock()
	if _, ok := h.clients[c]; ok {
		close(c.sendCh)
		delete(h.clients, c)
	}
}

// Broadcast 广播消息给所有客户端
// 参数：
//   - message []byte: 要广播的消息
//
// 返回值：
//   - 无
func (h *hub) broadcast(message []byte) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	for c := range h.clients {
		select {
		case c.sendCh <- message:
		default:
			// 如果队列满了，踢掉 client，避免阻塞
			go h.RemoveClient(c)
		}
	}
}

// StartBroadcast 启动广播 goroutine
// 这个 goroutine 负责从 logQueue 中读取消息并广播给所有客户端
// 参数：
//   - 无
//
// 返回值：
//   - 无
func (h *hub) startBroadcast() {
	go func() {
		for msg := range logQueue {
			h.broadcast(msg)
		}
	}()
}

// webSocketWriter 实现 io.Writer 接口，用于将日志写入 WebSocket
type webSocketWriter struct{}

// Write 实现 io.Writer 接口的 Write 方法
// 参数：
//   - p []byte: 要写入的日志数据
//
// 返回值：
//   - n int: 写入的字节数
//   - err error: 错误信息
func (w *webSocketWriter) Write(p []byte) (n int, err error) {
	msg := FormatLogMessage(p)
	logQueue <- []byte(msg)
	return len(p), nil
}

// NewWebSocketWriter 创建一个新的 WebSocketWriter
// 参数：
//   - 无
//
// 返回值：
//   - *webSocketWriter: 新创建的 WebSocketWriter
func NewWebSocketWriter() *webSocketWriter {
	if !variable.IsInitialized.Load() {
		LogHub.startBroadcast()
		return &webSocketWriter{}
	}
	return nil
}
