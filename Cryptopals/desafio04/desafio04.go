package desafio04

/*
Desafio 04 - Detectando Single-Byte XOR
Neste desafio, você precisa detectar qual linha em um arquivo
contém uma cifra que foi cifrada usando uma cifra XOR de byte único. Para isso,
você deve ler cada linha do arquivo, decodificar a string hexadecimal e tentar
decifrar o texto usando todas as possíveis chaves de 1 byte (0-255). Para cada
linha, você deve calcular a pontuação do texto decifrado e manter o controle
da melhor pontuação encontrada até o momento. No final, você deve retornar o texto decifrado
com a melhor pontuação, juntamente com a chave usada e o número da linha.
O código reutiliza a lógica de pontuação do desafio 3.
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
