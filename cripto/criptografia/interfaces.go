/*
Interfaces que definem contratos das estratégias criptográficas empregadas no
projeto, permitindo trocar algoritmos em tempo de execução sem afetar o
orquestrador de serviços.
*/
package criptografia

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
)

type IEstrategiaChave interface {
	GerarChavePrivada(tamanho int) (ParDeChaves, error)
}

type IEstrategiaCertificado interface {
	GerarCertificadoAutoassinado(chavePrivada *rsa.PrivateKey,
		sujeito *pkix.Name, validadeEmAnos int) (Certificado, error)
	GerarCertificadoAssinadoPorAC(
		chavePrivadaSujeito *rsa.PrivateKey,
		sujeito pkix.Name, validadeEmAnos int,
		certPai *x509.Certificate,
		chavePrivadaPai *rsa.PrivateKey) (Certificado, error)
}

type IEstrategiaResumo interface {
	Resumir(data []byte) []byte
	HashFunc() crypto.Hash
}

type IEstrategiaAssinatura interface {
	Assinar(resumo []byte,
		privKey *rsa.PrivateKey,
		hashFunc crypto.Hash) ([]byte, error)
	VerificarAssinatura(resumo []byte,
		assinatura []byte, pubKey *rsa.PublicKey,
		hashFunc crypto.Hash) error
}
