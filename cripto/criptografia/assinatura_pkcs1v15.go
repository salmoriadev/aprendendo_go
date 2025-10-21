package criptografia

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

type estrategiaAssinaturaPkcs1v15 struct{}

func (e *estrategiaAssinaturaPkcs1v15) Assinar(resumo []byte,
	chavePrivada *rsa.PrivateKey, hashFunc crypto.Hash) ([]byte, error) {

	assinatura, err := rsa.SignPKCS1v15(rand.Reader,
		chavePrivada, hashFunc, resumo)

	if err != nil {
		return nil, fmt.Errorf(
			"erro ao assinar o resumo (PKCS1v15): %w", err)
	}
	return assinatura, nil
}

func (e *estrategiaAssinaturaPkcs1v15) VerificarAssinatura(resumo,
	assinatura []byte, chavePublica *rsa.PublicKey,
	hashFunc crypto.Hash) error {

	err := rsa.VerifyPKCS1v15(chavePublica,
		hashFunc, resumo, assinatura)

	if err != nil {
		return fmt.Errorf(
			"a verificação da assinatura (PKCS1v15) falhou: %w", err)
	}
	return nil
}

func NovaEstrategiaAssinaturaPkcs1v15() EstrategiaAssinatura {
	return &estrategiaAssinaturaPkcs1v15{}
}
