package criptografia

import (
	"crypto/x509"
	"encoding/pem"
)

type Certificado struct {
	CertificadoBytes []byte
	Certificado      *x509.Certificate
}

func CertificadoParaPEM(certificado *Certificado) []byte {
	blocoPEM := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certificado.CertificadoBytes,
	}
	return pem.EncodeToMemory(blocoPEM)
}
