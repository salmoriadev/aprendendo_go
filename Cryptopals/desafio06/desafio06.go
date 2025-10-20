package desafio06

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

func DistanciaDeBits(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("listas devem ter o mesmo tamanho")
	}

	distancia := 0
	for i := 0; i < len(a); i++ {
		xorByte := a[i] ^ b[i]
		distancia += bits.OnesCount8(xorByte)
	}
	return distancia, nil
}

func AcharTamanhoChave(textoCifrado []byte, tamanhoMinimo, tamanhoMaximo, numeroBlocos int) []PalpiteTamanhoDaChave {
	var palpites []PalpiteTamanhoDaChave

	for tamanhoChave := tamanhoMinimo; tamanhoChave <= tamanhoMaximo; tamanhoChave++ {
		if (numeroBlocos * tamanhoChave) > len(textoCifrado) {
			break
		}

		var distanciaTotal float64
		var comparacao int

		for i := 0; i < numeroBlocos-1; i++ {
			for j := i + 1; j < numeroBlocos; j++ {
				bloco1 := textoCifrado[i*tamanhoChave : (i+1)*tamanhoChave]
				bloco2 := textoCifrado[j*tamanhoChave : (j+1)*tamanhoChave]

				distancia, err := DistanciaDeBits(bloco1, bloco2)
				if err != nil {
					continue
				}
				distanciaTotal += float64(distancia) / float64(tamanhoChave)
				comparacao++
			}
		}

		if comparacao > 0 {
			distanciaMedia := distanciaTotal / float64(comparacao)
			palpites = append(palpites, PalpiteTamanhoDaChave{Tamanho: tamanhoChave, Distancia: distanciaMedia})
		}
	}

	sort.Slice(palpites, func(i, j int) bool {
		return palpites[i].Distancia < palpites[j].Distancia
	})

	return palpites
}

func ResolverSingleByteXOR(textoCifrado []byte) (chave byte, textoPlano []byte, pontuacao float64) {
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
			pontuacaoAtual += float64(desafio03.FrequenciaLetrasEmIngles[unicode.ToLower(rune(decryptedByte))])
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
		blocosTranspostos[indexBloco] = append(blocosTranspostos[indexBloco], byteCifrado)
	}

	for i := 0; i < tamanhoChave; i++ {
		byteDaChave, _, _ := ResolverSingleByteXOR(blocosTranspostos[i])
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
		log.Fatalf("Implementacao de hamming distance esta incorreta! Esperado 37, retornado %d. Erro: %v", distanciaTeste, err)
	}

	textoCifradoBase64, err := os.ReadFile("desafio06/6.txt")
	if err != nil {
		return "", "", fmt.Errorf("falha ao ler o arquivo 6.txt: %w", err)
	}

	textoCifrado, err := base64.StdEncoding.DecodeString(string(textoCifradoBase64))
	if err != nil {
		return "", "", fmt.Errorf("falha ao decodificar base64: %w", err)
	}
	tamanhoMinimo := 2
	tamanhoMaximo := 40
	numeroBlocos := 4

	PalpiteTamanhoDaChave := AcharTamanhoChave(textoCifrado, tamanhoMinimo, tamanhoMaximo, numeroBlocos)

	melhorTamanhoChave := PalpiteTamanhoDaChave[0].Tamanho
	chave := AchaChaveRepetida(textoCifrado, melhorTamanhoChave)
	plaintext := XorComChaveRepetida(textoCifrado, chave)

	return string(plaintext), string(chave), nil
}
