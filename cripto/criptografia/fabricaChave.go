/*
Implementação da estratégia de geração de chaves RSA, isolando a criação e
expondo o contrato definido em interfaces para permitir troca futura por outros
algoritmos.
*/
package criptografia

import (
	"crypto/rand"
	"crypto/rsa"
)

type estrategiaChaveRSA struct{}

// NovoParDeChaves cria uma estrutura inicializada para receber as chaves RSA geradas.
func (f *estrategiaChaveRSA) NovoParDeChaves() ParDeChaves {
	return ParDeChaves{}
}

// GerarChavePrivada cria um novo par RSA com o tamanho solicitado,
// preenchendo a estrutura de retorno.
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

// NovaEstrategiaChaveRSA expõe a implementação de geração de chaves RSA via interface.
func NovaEstrategiaChaveRSA() IEstrategiaChave {
	return &estrategiaChaveRSA{}
}
