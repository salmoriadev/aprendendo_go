package desafio03

/*
Quebra de Single-Byte XOR (Desafio 03)

Exaustivamente testa as 256 chaves possíveis, pontua os textos decifrados com
um mapa de frequências em inglês e retorna o melhor palpite. Os utilitários são
reaproveitados pelos desafios 4 e 6.
*/

import (
	"encoding/hex"
	"fmt"
	"unicode"
)

var FrequenciaLetrasEmIngles = map[rune]int{
	' ': 1800,
	'a': 757, 'b': 184,
	'c': 409, 'd': 338,
	'e': 1151, 'f': 123,
	'g': 270, 'h': 232,
	'i': 901, 'j': 16,
	'k': 85, 'l': 531,
	'm': 284, 'n': 685,
	'o': 659, 'p': 294,
	'q': 16, 'r': 707,
	's': 952, 't': 668,
	'u': 327, 'v': 98,
	'w': 74, 'x': 29,
	'y': 163, 'z': 47,
}

type ResultadoDecifrado struct {
	Texto     string
	Pontuacao int
	Chave     byte
}

/*
Desafio03 resolve a cifra single-byte disponibilizada pelo enunciado, retornando
texto, pontuação e chave.
*/
func Desafio03() (ResultadoDecifrado, error) {
	hexCifra := "1b37373331363f78151b7f" +
		"2b783431333d78397828372d363c78373e783a393b3736"
	resultado, err := AcharChave(hexCifra)
	if err != nil {
		return ResultadoDecifrado{}, err
	}

	return resultado, nil
}

/*
AcharChave percorre as 256 possibilidades e usa PontuacaoTexto para definir o
melhor resultado.
*/
func AcharChave(entradaHex string) (ResultadoDecifrado, error) {
	entradaHexBytes, err := hex.DecodeString(entradaHex)
	if err != nil {
		return ResultadoDecifrado{}, fmt.Errorf(
			"falha ao decodificar a string hex: %w", err)
	}

	var melhorResultado ResultadoDecifrado

	for chave := 0; chave < 256; chave++ {
		textoEmBytes := make([]byte, len(entradaHexBytes))
		for i := 0; i < len(entradaHexBytes); i++ {
			textoEmBytes[i] = entradaHexBytes[i] ^ byte(chave)
		}

		pontuacaoAtual := PontuacaoTexto(textoEmBytes)

		if chave == 0 || pontuacaoAtual > melhorResultado.Pontuacao {
			melhorResultado = ResultadoDecifrado{
				Texto:     string(textoEmBytes),
				Pontuacao: pontuacaoAtual,
				Chave:     byte(chave),
			}
		}
	}
	return melhorResultado, nil
}

/*
PontuacaoTexto soma frequências aproximadas de inglês e penaliza caracteres
não imprimíveis.
*/
func PontuacaoTexto(textoBytes []byte) int {
	pontuacao := 0
	for _, bytes := range textoBytes {
		char := rune(bytes)
		transformarMinusculo := unicode.ToLower(char)

		pontuacao += FrequenciaLetrasEmIngles[transformarMinusculo]

		// Tiro pontos de textos com caracteres não usados na lingua inglesa
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) &&
			!unicode.IsSpace(char) && char != '\'' && char != ',' && char != '.' {
			pontuacao -= 5
		}
	}
	return pontuacao
}
