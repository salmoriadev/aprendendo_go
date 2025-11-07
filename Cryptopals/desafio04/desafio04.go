package desafio04

/*
Detecção de Single-Byte XOR (Desafio 04)

Varre a lista de hex strings fornecida pelo desafio 4, decifra cada uma com a
rotina do desafio 3 e devolve a linha mais provável. Complementa o Set 1 sem
se preocupar com leitura de arquivos externos.
*/

import (
	"cryptopals/desafio03"
	"encoding/hex"
	"fmt"
	"log"
)

type Decifragem struct {
	desafio03.ResultadoDecifrado
	NumeroLinha int
}

/*
Desafio04 percorre a lista de hex strings e devolve a melhor candidata de
cifra single-byte.
*/
func Desafio04() (Decifragem, error) {
	var melhorResultado Decifragem

	for i, hexLinha := range hexLista {
		linhaAtual := i + 1
		resultadoLinha, err := DecifraByteCifra(
			hexLinha, linhaAtual)

		if err != nil {
			log.Printf("Erro ao processar "+
				"linha %d: %v", linhaAtual, err)
			continue
		}

		if i == 0 || resultadoLinha.Pontuacao > melhorResultado.Pontuacao {
			melhorResultado = resultadoLinha
		}
	}

	if len(hexLista) > 0 {
		return melhorResultado, nil
	}
	return Decifragem{}, fmt.Errorf("nenhuma linha foi processada com sucesso")
}

/*
DecifraByteCifra isola a lógica de tentativa das 256 chaves para uma linha.
*/
func DecifraByteCifra(linhaHex string, numeroLinha int) (Decifragem, error) {
	cifra, err := hex.DecodeString(linhaHex)
	if err != nil {
		return Decifragem{}, fmt.Errorf("falha ao decodificar hex: %w", err)
	}

	var melhorDecifragemLinha Decifragem

	for chave := 0; chave < 256; chave++ {
		textoBytes := make([]byte, len(cifra))
		for i := 0; i < len(cifra); i++ {
			textoBytes[i] = cifra[i] ^ byte(chave)
		}

		pontuacaoAtual := desafio03.PontuacaoTexto(textoBytes)

		if chave == 0 || pontuacaoAtual > melhorDecifragemLinha.Pontuacao {
			melhorDecifragemLinha = Decifragem{
				ResultadoDecifrado: desafio03.ResultadoDecifrado{
					Texto:     string(textoBytes),
					Pontuacao: pontuacaoAtual,
					Chave:     byte(chave),
				},
				NumeroLinha: numeroLinha,
			}
		}
	}
	return melhorDecifragemLinha, nil
}
