package criptografia

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

type ParDeChaves struct {
	ChavePublica *rsa.PublicKey
	ChavePrivada *rsa.PrivateKey
}

func NovoParDeChaves() ParDeChaves {
	return ParDeChaves{}
}

func GerarChavePrivada(tamanho int) (ParDeChaves, error) {
	chave := NovoParDeChaves()
	var privada, err = rsa.GenerateKey(rand.Reader, tamanho)
	if err != nil {
		return chave, err
	}
	chave.ChavePrivada = privada
	chave.ChavePublica = &privada.PublicKey
	return chave, nil
}

func ChavePrivadaParaPEM(chavePrivada *rsa.PrivateKey) []byte {
	blocoPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(chavePrivada),
	}
	return pem.EncodeToMemory(blocoPEM)
}

func ChavePublicaParaPEM(chavePublica *rsa.PublicKey) []byte {
	blocoPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(chavePublica),
	}
	return pem.EncodeToMemory(blocoPEM)
}
