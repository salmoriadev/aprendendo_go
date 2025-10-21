package criptografia

import (
	"crypto/rand"
	"crypto/rsa"
)

type estrategiaChaveRSA struct{}

func (f *estrategiaChaveRSA) NovoParDeChaves() ParDeChaves {
	return ParDeChaves{}
}

func (f *estrategiaChaveRSA) GerarChavePrivada(
	tamanho int) (ParDeChaves, error) {
	chave := f.NovoParDeChaves()
	privada, err := rsa.GenerateKey(rand.Reader, tamanho)
	if err != nil {
		return chave, err
	}
	chave.ChavePrivada = privada
	chave.ChavePublica = &privada.PublicKey
	return chave, nil
}

func NovaEstrategiaChaveRSA() EstrategiaChave {
	return &estrategiaChaveRSA{}
}
