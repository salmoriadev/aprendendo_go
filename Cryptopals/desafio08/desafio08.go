package desafio08

/*
Desafio 08 - Detectando AES em modo ECB
Neste desafio, você precisa detectar qual linha em um arquivo contém
um texto cifrado usando o algoritmo AES em modo ECB. O modo ECB é
um modo de operação de cifra simétrica que divide o texto em blocos e
cifra cada bloco independentemente usando a mesma chave. Isso traz
vulnerabilidades, como a possibilidade de blocos repetidos no texto cifrado
indicar padrões no texto original. Esse é a base para decifrar o desafio.
*/

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

const tamanhoDoBloco = 16

type ResultadoECB struct {
	NumeroLinha   int
	MaxRepeticoes int
	Cifra         string
}

func AcharECB(dadosCifrados []byte) int {
	blocos := make(map[string]int)
	for i := 0; i < len(dadosCifrados); i += tamanhoDoBloco {
		bloco := string(dadosCifrados[i : i+tamanhoDoBloco])
		blocos[bloco]++
	}

	numRepeticoes := 0
	for _, count := range blocos {
		if count > 1 {
			numRepeticoes += (count - 1)
		}
	}
	return numRepeticoes
}

func Desafio08() (ResultadoECB, error) {
	arquivo, err := os.Open("desafio08/8.txt")
	if err != nil {
		return ResultadoECB{}, fmt.Errorf(
			"falha ao ler o arquivo 8.txt: %w", err)
	}
	defer arquivo.Close()
	primeiraLinha := 0
	resultado := ResultadoECB{
		NumeroLinha:   primeiraLinha - 1,
		MaxRepeticoes: primeiraLinha,
	}
	numeroLinhaAtual := primeiraLinha

	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		numeroLinhaAtual++
		linhaHex := scanner.Text()

		dadosCifrados, err := hex.DecodeString(linhaHex)
		if err != nil {
			continue
		}

		numRepeticoes := AcharECB(dadosCifrados)

		if numRepeticoes > resultado.MaxRepeticoes {
			resultado.MaxRepeticoes = numRepeticoes
			resultado.NumeroLinha = numeroLinhaAtual
			resultado.Cifra = linhaHex
		}
	}

	if err := scanner.Err(); err != nil {
		return ResultadoECB{}, fmt.Errorf(
			"erro ao escanear o arquivo: %w", err)
	}

	if resultado.NumeroLinha == -1 {
		return ResultadoECB{}, fmt.Errorf(
			"nenhuma linha com blocos repetidos foi encontrada")
	}

	return resultado, nil
}
