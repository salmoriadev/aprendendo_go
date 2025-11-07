package desafio06

/*
Quebra de XOR com Chave Repetida (Desafio 06)

Resolve o clássico desafio do Set 1: estimar o tamanho da chave via distância
de Hamming, transpor os blocos e tratar cada coluna como um single-byte XOR
reutilizando a pontuação do desafio 3. O objetivo é exclusivamente decifrar o
arquivo `desafio06/6.txt` sem generalizações extras.
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

/*
DistanciaDeBits calcula a distância de Hamming entre dois slices de mesmo
tamanho.
*/
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

/*
AcharTamanhoChave testa tamanhos possíveis avaliando a distância de Hamming
normalizada entre blocos consecutivos. Para o Set 1, ignoramos tamanhos que não
coubem na amostragem solicitada ao invés de abortar a execução inteira.
*/
func AcharTamanhoChave(textoCifrado []byte, tamanhoMinimo,
	tamanhoMaximo, numeroBlocos int) ([]PalpiteTamanhoDaChave, error) {
	var palpites []PalpiteTamanhoDaChave

	for tamanhoChave := tamanhoMinimo; tamanhoChave <= tamanhoMaximo; tamanhoChave++ {
		if (numeroBlocos * tamanhoChave) > len(textoCifrado) {
			continue
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

/*
ResolverSingleByteXOR recicla a lógica do desafio 3 para achar a melhor chave
de um XOR de byte único.
*/
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

/*
AchaChaveRepetida transpõe os blocos e resolve cada coluna como single-byte.
*/
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

/*
XorComChaveRepetida aplica a chave repetidamente sobre o buffer de entrada.
*/
func XorComChaveRepetida(entrada, chave []byte) []byte {
	saida := make([]byte, len(entrada))
	for i := 0; i < len(entrada); i++ {
		byteDaChave := chave[i%len(chave)]

		saida[i] = entrada[i] ^ byteDaChave
	}
	return saida
}

/*
Desafio06 executa ponta a ponta a solução do Set 1 para o arquivo `6.txt`.
*/
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
	if len(PalpiteTamanhoDaChave) == 0 {
		return "", "", fmt.Errorf(
			"nenhum tamanho de chave válido encontrado para o ciphertext")
	}

	melhorTamanhoChave := PalpiteTamanhoDaChave[0].Tamanho
	chave := AchaChaveRepetida(textoCifrado, melhorTamanhoChave)
	plaintext := XorComChaveRepetida(textoCifrado, chave)

	return string(plaintext), string(chave), nil
}
