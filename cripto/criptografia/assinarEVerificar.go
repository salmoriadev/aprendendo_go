package criptografia

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func AssinarResumo(resumo []byte, chavePrivada *rsa.PrivateKey) ([]byte, error) {
	assinatura, err := chavePrivada.Sign(rand.Reader, resumo, crypto.SHA256)
	if err != nil {
		return nil, fmt.Errorf("erro ao assinar o resumo: %w", err)
	}
	return assinatura, nil
}

func VerificarAssinatura(resumo, assinatura []byte, chavePublica *rsa.PublicKey) error {
	err := rsa.VerifyPKCS1v15(chavePublica, crypto.SHA256, resumo, assinatura)
	if err != nil {
		return fmt.Errorf("a verificação da assinatura falhou: %w", err)
	}
	return nil
}
