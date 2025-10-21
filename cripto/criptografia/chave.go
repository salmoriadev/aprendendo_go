package criptografia

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

type ParDeChaves struct {
	ChavePublica *rsa.PublicKey
	ChavePrivada *rsa.PrivateKey
}

func ChavePrivadaParaPEM(chavePrivada *rsa.PrivateKey) []byte {
	blocoPEM := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(chavePrivada),
	}
	return pem.EncodeToMemory(blocoPEM)
}

func ChavePublicaParaPEM(chavePublica *rsa.PublicKey) []byte {
	blocoPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(chavePublica),
	}
	return pem.EncodeToMemory(blocoPEM)
}
