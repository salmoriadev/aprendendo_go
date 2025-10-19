package criptografia

import (
	"crypto/rsa"
)

type EstrategiaResumo interface {
	Resumir(data []byte) []byte
}

type EstrategiaAssinatura interface {
	Assinar(resumo []byte, privKey *rsa.PrivateKey) ([]byte, error)
	VerificarAssinatura(resumo []byte,
		assinatura []byte, pubKey *rsa.PublicKey) error
}
