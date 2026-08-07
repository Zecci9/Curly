package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func main() {

	// 生成私钥
	privateKey, err := rsa.GenerateKey(
		rand.Reader,
		2048,
	)

	if err != nil {
		panic(err)
	}

	// 证书信息
	template := x509.Certificate{

		SerialNumber: big.NewInt(1),

		Subject: pkix.Name{
			CommonName: "Curly Local",
		},

		NotBefore: time.Now(),

		NotAfter: time.Now().AddDate(
			1,
			0,
			0,
		),

		KeyUsage: x509.KeyUsageKeyEncipherment |
			x509.KeyUsageDigitalSignature,

		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
		},

		IPAddresses: []net.IP{
			net.ParseIP("127.0.0.1"),
			net.ParseIP("192.168.8.48"),
		},

		DNSNames: []string{
			"localhost",
		},
	}

	certBytes, err := x509.CreateCertificate(
		rand.Reader,
		&template,
		&template,
		&privateKey.PublicKey,
		privateKey,
	)

	if err != nil {
		panic(err)
	}

	// 保存证书
	certFile, _ := os.Create("server.crt")

	pem.Encode(
		certFile,
		&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: certBytes,
		},
	)

	certFile.Close()

	// 保存私钥
	keyFile, _ := os.Create("server.key")

	pem.Encode(
		keyFile,
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)

	keyFile.Close()

}
