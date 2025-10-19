package criptografia

import (
	"crypto/sha256"
)

type ResumoSha256 struct{}

func (e *ResumoSha256) Resumir(dados []byte) []byte {
	hash := sha256.Sum256(dados)
	return hash[:]
}
