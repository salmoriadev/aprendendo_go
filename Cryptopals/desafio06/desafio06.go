package desafio06

/*
Desafio 06 - Quebrando o XOR com chave repetida
Neste desafio, você precisa decifrar um texto que foi cifrado usando
uma cifra XOR com chave repetida com o tamanho da chave desconhecido.
Me baseei na explicação do desafio para implementar a solução, acredito que pela
maior dificuldade dele existiu uma necessidade de uma explicação maior.
Primeiramente você deve determinar o tamanho da chave usando a distância de Hamming;
Como blocos em inglês têm padrões semelhantes, você pode usar a distância
de Hamming (ou distância de bits) para estimar o tamanho da chave.
Após isso vai dividir o texto cifrado em blocos do tamanho da chave.
Vai transpor os blocos para que cada bloco contenha os bytes correspondentes.
Vai resolver cada bloco como um problema de XOR de byte único.
E por último, reunir os resultados para obter o texto decifrado.
O código reutiliza o mapa de pontuação e uma função do desafio 3.
*/

import (
	"cryptopals/desafio03"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"math/bits"
	"os"
	"sort"
	"unicode"
)

type PalpiteTamanhoDaChave struct {
	Tamanho   int
	Distancia float64
}

func DistanciaDeBits(lista1, lista2 []byte) (int, error) {
	if len(lista1) != len(lista2) {
		return 0, errors.New("listas devem ter o mesmo tamanho")
	}

	distancia := 0
	for i := 0; i < len(lista1); i++ {
		xorByte := lista1[i] ^ lista2[i]
		distancia += bits.OnesCount8(uint8(xorByte))
	}
	return distancia, nil
}

func AcharTamanhoChave(textoCifrado []byte, tamanhoMinimo,
	tamanhoMaximo, numeroBlocos int) ([]PalpiteTamanhoDaChave, error) {
	var palpites []PalpiteTamanhoDaChave

	for tamanhoChave := tamanhoMinimo; tamanhoChave <= tamanhoMaximo; tamanhoChave++ {
		if (numeroBlocos * tamanhoChave) > len(textoCifrado) {
			return nil, fmt.Errorf(
				"texto cifrado muito curto para o tamanho"+
					" da chave %d e número de blocos %d",
				tamanhoChave, numeroBlocos)
		}

		var distanciaTotal float64
		var comparacao int

		for i := 0; i < numeroBlocos-1; i++ {
			for j := i + 1; j < numeroBlocos; j++ {
				bloco1 := textoCifrado[i*tamanhoChave : (i+1)*tamanhoChave]
				bloco2 := textoCifrado[j*tamanhoChave : (j+1)*tamanhoChave]

				distancia, err := DistanciaDeBits(bloco1, bloco2)
				if err != nil {
					return nil, fmt.Errorf(
						"erro ao calcular a distância de bits: %w", err)
				}
				distanciaTotal += float64(distancia) / float64(tamanhoChave)
				comparacao++
			}
		}

		if comparacao > 0 {
			distanciaMedia := distanciaTotal / float64(comparacao)
			palpites = append(palpites, PalpiteTamanhoDaChave{
				Tamanho:   tamanhoChave,
				Distancia: distanciaMedia,
			})
		}
	}

	sort.Slice(palpites, func(i, j int) bool {
		return palpites[i].Distancia < palpites[j].Distancia
	})

	return palpites, nil
}

func ResolverSingleByteXOR(textoCifrado []byte) (chave byte,
	textoPlano []byte, pontuacao float64) {
	var maxPontuacao float64 = -1.0
	var melhorChave byte
	var melhorTextoPlano []byte

	for k := 0; k < 256; k++ {
		palpiteChave := byte(k)
		textoPlanoAtual := make([]byte, len(textoCifrado))
		var pontuacaoAtual float64

		for i := 0; i < len(textoCifrado); i++ {
			decryptedByte := textoCifrado[i] ^ palpiteChave
			textoPlanoAtual[i] = decryptedByte
			pontuacaoAtual += float64(
				desafio03.FrequenciaLetrasEmIngles[unicode.ToLower(
					rune(decryptedByte))])
		}

		if pontuacaoAtual > maxPontuacao {
			maxPontuacao = pontuacaoAtual
			melhorChave = palpiteChave
			melhorTextoPlano = textoPlanoAtual
		}
	}
	return melhorChave, melhorTextoPlano, maxPontuacao
}

func AchaChaveRepetida(textoCifrado []byte, tamanhoChave int) []byte {
	chave := make([]byte, tamanhoChave)
	blocosTranspostos := make([][]byte, tamanhoChave)
	for i := 0; i < tamanhoChave; i++ {
		blocosTranspostos[i] = make([]byte, 0)
	}

	for i, byteCifrado := range textoCifrado {
		indexBloco := i % tamanhoChave
		blocosTranspostos[indexBloco] = append(
			blocosTranspostos[indexBloco], byteCifrado)
	}

	for i := 0; i < tamanhoChave; i++ {
		byteDaChave, _, _ := ResolverSingleByteXOR(
			blocosTranspostos[i])
		chave[i] = byteDaChave
	}

	return chave
}

func XorComChaveRepetida(entrada, chave []byte) []byte {
	saida := make([]byte, len(entrada))
	for i := 0; i < len(entrada); i++ {
		byteDaChave := chave[i%len(chave)]

		saida[i] = entrada[i] ^ byteDaChave
	}
	return saida
}

func Desafio06() (string, string, error) {

	frase01 := []byte("this is a test")
	frase02 := []byte("wokka wokka!!!")
	distanciaTeste, err := DistanciaDeBits(frase01, frase02)
	if err != nil || distanciaTeste != 37 {
		log.Fatalf(
			"Implementacao de hamming distance esta incorreta!"+
				" Esperado 37, retornado %d. Erro: %v",
			distanciaTeste, err)
	}

	textoCifradoBase64, err := os.ReadFile("desafio06/6.txt")
	if err != nil {
		return "", "", fmt.Errorf("falha ao ler o arquivo 6.txt: %w", err)
	}

	textoCifrado, err := base64.StdEncoding.DecodeString(
		string(textoCifradoBase64))
	if err != nil {
		return "", "", fmt.Errorf("falha ao decodificar base64: %w", err)
	}
	tamanhoMinimo := 2
	tamanhoMaximo := 40
	numeroBlocos := 4

	PalpiteTamanhoDaChave, err := AcharTamanhoChave(
		textoCifrado, tamanhoMinimo,
		tamanhoMaximo, numeroBlocos)
	if err != nil {
		return "", "", fmt.Errorf(
			"falha ao achar o tamanho da chave: %w", err)
	}

	melhorTamanhoChave := PalpiteTamanhoDaChave[0].Tamanho
	chave := AchaChaveRepetida(textoCifrado, melhorTamanhoChave)
	plaintext := XorComChaveRepetida(textoCifrado, chave)

	return string(plaintext), string(chave), nil
}
