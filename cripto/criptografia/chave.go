/*
Estruturas e utilitários relacionados a pares de chaves RSA, focados em manter
os dados em memória e convertê-los para formato PEM sem qualquer dependência de
I/O.
*/
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

// ChavePrivadaParaPEM serializa uma chave privada RSA no formato PKCS#1,
// retornando bytes prontos para salvar.
func ChavePrivadaParaPEM(chavePrivada *rsa.PrivateKey) []byte {
	blocoPEM := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(chavePrivada),
	}
	return pem.EncodeToMemory(blocoPEM)
}

// ChavePublicaParaPEM serializa a chave pública RSA no formato PKCS#1 para
// facilitar sua distribuição.
func ChavePublicaParaPEM(chavePublica *rsa.PublicKey) []byte {
	blocoPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(chavePublica),
	}
	return pem.EncodeToMemory(blocoPEM)
}
