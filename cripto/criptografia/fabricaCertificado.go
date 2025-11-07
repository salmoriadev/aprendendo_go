/*
Fábricas de certificados X.509 responsáveis por gerar tanto a autoridade quanto
os certificados finais, mantendo o foco em manipular objetos puros sem acesso a
disco.
*/
package criptografia

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"time"
)

// Implementacoes concretas das estrategias de certificado X.509.
// Mantem as regras proximas dos artefatos criptograficos e livres de I/O.
type estrategiaCertificado struct{}

// novoCertificado inicializa a estrutura de retorno utilizada ao criar certificados.
func (f *estrategiaCertificado) novoCertificado() Certificado {
	return Certificado{}
}

// GerarCertificadoAutoassinado constroi um certificado raiz com uso e extensoes minimas.
func (f *estrategiaCertificado) GerarCertificadoAutoassinado(
	chavePrivada *rsa.PrivateKey,
	sujeito *pkix.Name, validadeEmAnos int) (Certificado, error) {

	cert := f.novoCertificado()
	inicioPrazo := time.Now()
	validade := inicioPrazo.AddDate(validadeEmAnos, 0, 0)
	permissoesDaChave := x509.KeyUsageCertSign |
		x509.KeyUsageCRLSign
	propositosDaChave := []x509.ExtKeyUsage{
		x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth}

	serial, err := gerarNumeroSerial()
	if err != nil {
		return cert, fmt.Errorf("erro ao gerar número serial: %w", err)
	}

	cert.Certificado = &x509.Certificate{
		SerialNumber:          serial,
		Subject:               *sujeito,
		NotBefore:             inicioPrazo,
		NotAfter:              validade,
		KeyUsage:              permissoesDaChave,
		ExtKeyUsage:           propositosDaChave,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}
	err = anexarExtensoesDeIdentificacao(
		cert.Certificado, &chavePrivada.PublicKey, nil)
	if err != nil {
		return cert, err
	}

	cert.CertificadoBytes, err = x509.CreateCertificate(rand.Reader,
		cert.Certificado, cert.Certificado,
		&chavePrivada.PublicKey, chavePrivada)
	if err != nil {
		return cert, err
	}

	return cert, nil
}

// GerarCertificadoAssinadoPorAC emite um certificado final assinado pela autoridade informada.
func (f *estrategiaCertificado) GerarCertificadoAssinadoPorAC(
	chavePrivadaSujeito *rsa.PrivateKey,
	sujeito pkix.Name, validadeEmAnos int, certPai *x509.Certificate,
	chavePrivadaPai *rsa.PrivateKey) (Certificado, error) {

	cert := f.novoCertificado()
	inicioPrazo := time.Now()
	validade := inicioPrazo.AddDate(validadeEmAnos, 0, 0)
	permissoesDaChave := x509.KeyUsageKeyEncipherment |
		x509.KeyUsageDigitalSignature
	propositosDaChave := []x509.ExtKeyUsage{
		x509.ExtKeyUsageServerAuth}

	serial, err := gerarNumeroSerial()
	if err != nil {
		return cert, fmt.Errorf("erro ao gerar número serial: %w", err)
	}

	cert.Certificado = &x509.Certificate{
		SerialNumber:          serial,
		Subject:               sujeito,
		NotBefore:             inicioPrazo,
		NotAfter:              validade,
		KeyUsage:              permissoesDaChave,
		ExtKeyUsage:           propositosDaChave,
		BasicConstraintsValid: true,
		IsCA:                  false,
	}

	err = anexarExtensoesDeIdentificacao(
		cert.Certificado, &chavePrivadaSujeito.PublicKey, certPai)
	if err != nil {
		return cert, err
	}

	cert.CertificadoBytes, err = x509.CreateCertificate(
		rand.Reader, cert.Certificado, certPai,
		&chavePrivadaSujeito.PublicKey, chavePrivadaPai)
	if err != nil {
		return cert, err
	}
	return cert, nil
}

// NovaEstrategiaCertificado disponibiliza a implementação X.509 pronta para uso via interface.
func NovaEstrategiaCertificado() IEstrategiaCertificado {
	return &estrategiaCertificado{}
}

// gerarNumeroSerial fornece números aleatórios positivos,
// evitando colisões na emissão dos certificados.
func gerarNumeroSerial() (*big.Int, error) {
	limite := new(big.Int).Lsh(big.NewInt(1), 128)
	serial, err := rand.Int(rand.Reader, limite)
	if err != nil {
		return nil, err
	}
	if serial.Sign() == 0 {
		return gerarNumeroSerial()
	}
	return serial, nil
}

// anexarExtensoesDeIdentificacao adiciona SKI/AKI conforme RFC 5280.
// A derivacao usa SHA-1 por compatibilidade historica dos campos;
// o hash não é reutilizado fora disso.
func anexarExtensoesDeIdentificacao(
	cert *x509.Certificate, pubKey *rsa.PublicKey,
	certPai *x509.Certificate) error {
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return fmt.Errorf("erro ao serializar chave pública: %w", err)
	}
	resumo := sha1.Sum(pubKeyBytes)
	cert.SubjectKeyId = append([]byte(nil), resumo[:]...)

	if certPai != nil {
		if len(certPai.SubjectKeyId) == 0 {
			paiKeyBytes, err := x509.MarshalPKIXPublicKey(certPai.PublicKey)
			if err != nil {
				return fmt.Errorf("erro ao serializar chave pública da AC: %w", err)
			}
			paiResumo := sha1.Sum(paiKeyBytes)
			certPai.SubjectKeyId = append([]byte(nil), paiResumo[:]...)
		}
		cert.AuthorityKeyId = append([]byte(nil), certPai.SubjectKeyId...)
	}

	return nil
}
