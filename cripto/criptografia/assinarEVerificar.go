package criptografia

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

type AssinaturaPkcs1v15 struct{}

func (e *AssinaturaPkcs1v15) Assinar(resumo []byte, chavePrivada *rsa.PrivateKey) ([]byte, error) {
	assinatura, err := rsa.SignPKCS1v15(rand.Reader, chavePrivada, crypto.SHA256, resumo)
	if err != nil {
		return nil, fmt.Errorf("erro ao assinar o resumo (PKCS1v15): %w", err)
	}
	return assinatura, nil
}

func (e *AssinaturaPkcs1v15) VerificarAssinatura(resumo, assinatura []byte, chavePublica *rsa.PublicKey) error {
	err := rsa.VerifyPKCS1v15(chavePublica, crypto.SHA256, resumo, assinatura)
	if err != nil {
		return fmt.Errorf("a verificação da assinatura (PKCS1v15) falhou: %w", err)
	}
	return nil
}
