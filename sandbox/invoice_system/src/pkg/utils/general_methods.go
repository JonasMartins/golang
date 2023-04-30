package utils

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"os/exec"

	"google.golang.org/grpc/credentials"
)

func GenerateNewUUid() ([]byte, error) {

	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		return nil, err
	}
	return newUUID, nil
}

func LoadTLSCredentials() (credentials.TransportCredentials, error) {
	utils := New()
	path, err := utils.GetFilePath(&[]string{"cert", "ca-cert.pem"})
	if err != nil {
		return nil, err
	}
	pemServerCA, err := os.ReadFile(*path)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	cfg := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(cfg), nil
}
