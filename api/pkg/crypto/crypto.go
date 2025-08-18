package crypto

import (
	"crypto/rsa"
)

var (
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
)

func ParseRsaKeys(pubPEM, priPEM []byte) error {
	return nil
}
