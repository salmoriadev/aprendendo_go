/*
Serviços responsáveis por orquestrar a emissão de chaves e certificados,
conectando as estratégias criptográficas puras aos artefatos persistidos em
disco para demonstrar o fluxo completo de ICP.
*/
package servicos

import (
	"cripto/criptografia"
	"crypto/x509/pkix"
	"fmt"
	"path/filepath"
)

// Package servicos orquestra o fluxo de certificacao digital, conectando
// estrategias criptograficas puras com rotinas de I/O e permitindo testes isolados.

// DadosIdentificacaoCertificado descreve um sujeito em um certificado X.509.
// Os campos sao usados para compor o pkix.Name, mantendo a interface em portugues.
type DadosIdentificacaoCertificado struct {
	Organizacao string
	Pais        string
	Provincia   string
	Localidade  string
	NomeComum   string
}

// ParaPkixName transforma os dados em um pkix.Name consumido pelas estrategias.
func (d DadosIdentificacaoCertificado) ParaPkixName() pkix.Name {
	return pkix.Name{
		Organization: []string{d.Organizacao},
		Country:      []string{d.Pais},
		Province:     []string{d.Provincia},
		Locality:     []string{d.Localidade},
		CommonName:   d.NomeComum,
	}
}

// GerarParDeChaves cria um par RSA e salva as chaves em PEM no diretorio informado.
// Retorna o par em memoria para uso imediato, mantendo feedback textual sobre o fluxo.
func GerarParDeChaves(tamanhoChave int, caminhoArquivos string, nomeArquivo string,
	estrategia criptografia.IEstrategiaChave) (criptografia.ParDeChaves, error) {

	chaves, err := estrategia.GerarChavePrivada(tamanhoChave)
	if err != nil {
		return criptografia.ParDeChaves{}, fmt.Errorf(
			"falha ao gerar chave RSA (%s): %w", nomeArquivo, err)
	}

	fmt.Printf("Par de chaves RSA (%s) gerado com sucesso\n", nomeArquivo)

	dadosPEMPrivada := criptografia.ChavePrivadaParaPEM(chaves.ChavePrivada)
	caminhoChavePrivada := filepath.Join(caminhoArquivos,
		fmt.Sprintf("chave_%s_privada.pem", nomeArquivo))
	if err = escreverArquivo(caminhoChavePrivada, dadosPEMPrivada); err != nil {
		return criptografia.ParDeChaves{}, fmt.Errorf(
			"falha ao salvar chave privada (%s): %w", nomeArquivo, err)
	}

	dadosPEMPublica := criptografia.ChavePublicaParaPEM(chaves.ChavePublica)
	caminhoChavePublica := filepath.Join(caminhoArquivos,
		fmt.Sprintf("chave_%s_publica.pem", nomeArquivo))
	if err = escreverArquivo(caminhoChavePublica, dadosPEMPublica); err != nil {
		return criptografia.ParDeChaves{}, fmt.Errorf(
			"falha ao salvar chave pública (%s): %w", nomeArquivo, err)
	}

	return chaves, nil
}

// GerarCertificados emite o certificado da autoridade e do usuario, gravando-os em disco.
// Responsavel por conectar os sujeitos, validades e estrategias externas ao pacote criptografia.
func GerarCertificados(chaveAutoridade, chaveUsuario criptografia.ParDeChaves,
	validadeAutoridade int, validadeUsuario int, caminhoArquivos string,
	dadosAutoridade, dadosUsuario DadosIdentificacaoCertificado,
	estrategiaCertificado criptografia.IEstrategiaCertificado) error {

	distintivoAutoridade := dadosAutoridade.ParaPkixName()
	certificadoAutoridade, err := estrategiaCertificado.GerarCertificadoAutoassinado(
		chaveAutoridade.ChavePrivada, &distintivoAutoridade, validadeAutoridade)
	if err != nil {
		return fmt.Errorf("erro ao gerar certificado autoassinado da AC: %w", err)
	}

	certificadoAutoridadePEM := criptografia.CertificadoParaPEM(&certificadoAutoridade)
	caminhoCertificadoAutoridade := filepath.Join(caminhoArquivos, "certificado_autoridade.pem")
	if err = escreverArquivo(caminhoCertificadoAutoridade, certificadoAutoridadePEM); err != nil {
		return fmt.Errorf(
			"erro ao salvar certificado da AC (%s): %w", caminhoCertificadoAutoridade, err)
	}
	fmt.Println("Certificado da autoridade certificadora gerado com sucesso!")

	distintivoUsuario := dadosUsuario.ParaPkixName()
	certificadoUsuario, err := estrategiaCertificado.GerarCertificadoAssinadoPorAC(
		chaveUsuario.ChavePrivada, distintivoUsuario,
		validadeUsuario, certificadoAutoridade.Certificado, chaveAutoridade.ChavePrivada)
	if err != nil {
		return fmt.Errorf("erro ao gerar certificado do usuário: %w", err)
	}

	certificadoUsuarioPEM := criptografia.CertificadoParaPEM(&certificadoUsuario)
	caminhoCertificadoUsuario := filepath.Join(caminhoArquivos, "certificado_usuario.pem")
	if err = escreverArquivo(caminhoCertificadoUsuario, certificadoUsuarioPEM); err != nil {
		return fmt.Errorf(
			"erro ao salvar certificado do usuário (%s): %w",
			caminhoCertificadoUsuario, err)
	}
	fmt.Println("Certificado do usuário gerado e assinado com sucesso!")

	return nil
}
