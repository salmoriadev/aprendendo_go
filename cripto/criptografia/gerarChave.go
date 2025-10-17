package criptografia

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
)

type parDeChaves struct {
	ChavePublica *rsa.PublicKey
	ChavePrivada *rsa.PrivateKey
}

func NovaParDeChaves() parDeChaves {
	return parDeChaves{}
}

func GerarChavePrivada(tamanho int) (parDeChaves, error) {
	chave := NovaParDeChaves()
	var privada, err = rsa.GenerateKey(rand.Reader, tamanho)
	if err != nil {
		return chave, err
	}
	chave.ChavePrivada = privada
	chave.ChavePublica = &privada.PublicKey
	return chave, nil
}

func CodificarChaveParaBase64(chave []byte) string {
	return base64.StdEncoding.EncodeToString(chave)
}

func DecodificarChaveDeBase64(chaveBase64 string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(chaveBase64)
}

func EscreverChavePrivadaParaArquivoPEM(chavePrivada *rsa.PrivateKey, caminho string) error {
	arquivo, err := os.Create(caminho)
	if err != nil {
		return err
	}
	defer arquivo.Close()

	blocoPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(chavePrivada),
	}

	return pem.Encode(arquivo, blocoPEM)
}

func EscreverChavePublicaParaArquivoPEM(chavePublica *rsa.PublicKey, caminho string) error {
	arquivo, err := os.Create(caminho)
	if err != nil {
		return err
	}
	defer arquivo.Close()

	blocoPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(chavePublica),
	}

	return pem.Encode(arquivo, blocoPEM)
}
