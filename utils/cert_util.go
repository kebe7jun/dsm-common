package utils

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// Convert CA value to CA object, the CA value muse the pem.
func ParseToCertificate(value []byte) (*x509.Certificate, error) {
	roots := x509.NewCertPool()
	roots.AppendCertsFromPEM(value)
	block, _ := pem.Decode(value)
	if block == nil {
		// 解析异常
		return nil, errors.New("Can't analyse CA ")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	return cert, err
}
