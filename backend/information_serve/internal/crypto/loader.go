package crypto

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

type RSAKeyPair struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// LoadPrivateKey 加载私钥
func LoadPrivateKey(privatePath string) (*rsa.PrivateKey, error) {

	// 1️⃣ 读取私钥
	privBytes, err := os.ReadFile(privatePath)
	if err != nil {
		return nil, fmt.Errorf("read private key failed: %w", err)
	}

	privBlock, _ := pem.Decode(privBytes)
	if privBlock == nil {
		return nil, fmt.Errorf("invalid private pem")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse private key failed: %w", err)
	}

	return privateKey, nil
}

// 加载公钥
func LoadPublicKey(publicPath string) (*rsa.PublicKey, error) {

	// 2️⃣ 读取公钥
	pubBytes, err := os.ReadFile(publicPath)
	if err != nil {
		return nil, fmt.Errorf("read public key failed: %w", err)
	}

	pubBlock, _ := pem.Decode(pubBytes)
	if pubBlock == nil {
		return nil, fmt.Errorf("invalid public pem")
	}

	pubAny, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse public key failed: %w", err)
	}

	publicKey, ok := pubAny.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not rsa public key")
	}
	return publicKey, nil
}
