package desafio04

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

func Desafio04() {
	var melhorResultado Decifragem

	for i, hexLinha := range hexLista {
		linhaAtual := i + 1
		resultadoLinha, err := DecifraByteCifra(hexLinha, linhaAtual)

		if err != nil {
			log.Printf("Erro ao processar linha %d: %v", linhaAtual, err)
			continue
		}

		if i == 0 || resultadoLinha.Pontuacao > melhorResultado.Pontuacao {
			melhorResultado = resultadoLinha
		}
	}

	if len(hexLista) > 0 {
		fmt.Println("Linha: ", melhorResultado.NumeroLinha)
		fmt.Printf("Chave: %c\n", rune(melhorResultado.Chave))
		fmt.Println("Pontuação: ", melhorResultado.Pontuacao)
		fmt.Println("Texto: ", melhorResultado.Texto)
	}
}

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
