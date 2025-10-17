package main

import (
	"cripto/criptografia"
)

func main() {
	tamanhoChaves := 2048
	validadeCertAC := 10
	validadeCert := 1

	chaveAc := criptografia.ExecucaoChaves(tamanhoChaves, caminho)
	chaveCert := criptografia.ExecucaoChaves(tamanhoChaves, caminho)
	criptografia.ExecucaoCertificados(chaveAc, chaveCert, tamanhoChaves,
		validadeCertAC, validadeCert, caminho, "Meu Site",
		"BR", "Santa Catarina", "Florianopolis", "localhost")
}
