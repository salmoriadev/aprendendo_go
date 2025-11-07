/*
Estratégia de assinatura digital baseada em PKCS#1 v1.5, implementando o
contrato de assinatura e verificação utilizado pela camada de serviços.
*/
package criptografia

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

type estrategiaAssinaturaPkcs1v15 struct{}

// Assinar aplica PKCS#1 v1.5 sobre o resumo informado utilizando a chave privada.
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

// VerificarAssinatura compara a assinatura fornecida com o resumo gerado usando a chave pública.
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

// NovaEstrategiaAssinaturaPkcs1v15 expõe a estratégia de assinatura PKCS#1 v1.5 para injeção de dependência.
func NovaEstrategiaAssinaturaPkcs1v15() IEstrategiaAssinatura {
	return &estrategiaAssinaturaPkcs1v15{}
}
