package criptografia

import (
	"crypto"
	"crypto/sha256"
)

type estrategiaResumoSha256 struct{}

func (e *estrategiaResumoSha256) Resumir(dados []byte) []byte {
	hash := sha256.Sum256(dados)
	return hash[:]
}

func (e *estrategiaResumoSha256) HashFunc() crypto.Hash {
	return crypto.SHA256
}

func NovaEstrategiaResumoSha256() EstrategiaResumo {
	return &estrategiaResumoSha256{}
}
