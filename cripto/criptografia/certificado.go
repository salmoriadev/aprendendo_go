/*
Estruturas auxiliares para representar certificados X.509 em memória e
conversões para o formato PEM utilizados pelos serviços superiores.
*/
package criptografia

import (
	"crypto/x509"
	"encoding/pem"
)

type Certificado struct {
	CertificadoBytes []byte
	Certificado      *x509.Certificate
}

// CertificadoParaPEM converte o certificado X.509 em memória para o formato PEM textual.
func CertificadoParaPEM(certificado *Certificado) []byte {
	blocoPEM := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certificado.CertificadoBytes,
	}
	return pem.EncodeToMemory(blocoPEM)
}
