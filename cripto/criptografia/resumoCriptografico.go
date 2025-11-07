/*
Estratégia de resumo criptográfico SHA-256 exposta via interface, permitindo
que o restante do sistema trabalhe apenas com contratos abstraindo o algoritmo.
*/
package criptografia

import (
	"crypto"
	"crypto/sha256"
)

type estrategiaResumoSha256 struct{}

func (e *estrategiaResumoSha256) Resumir(dados []byte) []byte {
	// Resumir calcula o hash SHA-256 do conteúdo fornecido e devolve os bytes brutos.
	hash := sha256.Sum256(dados)
	return hash[:]
}

func (e *estrategiaResumoSha256) HashFunc() crypto.Hash {
	// HashFunc informa qual identificador de hash deve ser usado nas operações de assinatura.
	return crypto.SHA256
}

// NovaEstrategiaResumoSha256 disponibiliza a estratégia SHA-256 através da interface do domínio.
func NovaEstrategiaResumoSha256() IEstrategiaResumo {
	return &estrategiaResumoSha256{}
}
