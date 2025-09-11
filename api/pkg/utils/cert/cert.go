package cert

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/utils/file"
	"gen_gin_tpl/pkg/variable"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// GetCertificatePath 获取证书路径
// 参数：
//   - dataPath 数据路径
//
// 返回值：
//   - crtPath 证书路径
//   - keyPath 私钥路径
func GetCertificatePath(dataPath string) (crtPath, keyPath string) {
	crtPath = filepath.Join(dataPath, "CERTIFICATE.crt")
	keyPath = filepath.Join(dataPath, "EC_PRIVATE_KEY.key")
	return
}

// GenCertificateFile 生成证书
func GenCertificateFile(domainName, ip, dataPath string) {
	crtPath, keyPath := GetCertificatePath(dataPath)
	if file.IsFileExist(crtPath) && file.IsFileExist(keyPath) {
		return
	}
	// 生成私钥
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Error().Err(err).Msg("证书 P256 生成失败！")
		return
	}

	// 证书模板
	template := x509.Certificate{
		SerialNumber: big.NewInt(time.Now().UnixNano()),
		Subject: pkix.Name{
			Organization: []string{"Local Me"},
			CommonName:   "Local Me",
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(365 * 24 * time.Hour), // 有效期 1 年

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,

		DNSNames: []string{domainName},
		IPAddresses: []net.IP{
			net.ParseIP(ip), // 一定要确保不是 nil
		},
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		log.Error().Err(err).Msg("证书生成失败！")
		return
	}

	certOut, _ := os.Create(crtPath)
	_ = pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	_ = certOut.Close()

	// 写私钥文件
	keyOut, _ := os.Create(keyPath)
	b, _ := x509.MarshalECPrivateKey(priv)
	_ = pem.Encode(keyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: b})
	_ = keyOut.Close()

	log.Info().Msgf("✅ 证书生成完成: %s", crtPath)
	log.Info().Msgf("✅ 证书生成完成: %s", keyPath)
}

// InitRSAVariable 初始化RSA变量
func InitRSAVariable() {
	// 生成 2048 位 RSA 私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// 导出 PEM 格式的私钥
	privBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	variable.PrivatePEM = pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privBytes})

	// 从私钥生成公钥
	pubBytes, _ := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	variable.PublicPEM = pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pubBytes})

	privateLines := strings.Split(string(variable.PrivatePEM), "\n")
	variable.SessionKey = []byte(privateLines[len(privateLines)/2])
}
