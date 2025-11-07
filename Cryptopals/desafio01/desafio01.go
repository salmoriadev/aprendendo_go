package desafio01

/*
Conversão Hex -> Texto (Desafio 01)

Entrega a solução direta do primeiro desafio: transformar a string hex em ASCII
e devolver o conteúdo legível.
*/

import (
	"encoding/hex"
	"fmt"
)

/*
HexToString traduz a entrada hex em string. Em caso de erro, devolve string
vazia apenas para cumprir o fluxo dos desafios do Set 1.
*/
func HexToString(hexString string) string {
	decodedBytes, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println("Erro ao decodificar a string hexadecimal:", err)
		return ""
	}
	resultString := string(decodedBytes)
	return resultString
}

/*
Desafio01 é um invólucro que chama HexToString para o valor fornecido pelo
enunciado.
*/
func Desafio01() (string, error) {
	hexString := ("686974207468652062756c6c2773206" +
		"5796549276d206b696c6c696e6720796f757220627261696e206c696b" +
		"65206120706f69736f6e6f7573206d757368726f6f6d")
	resultString := HexToString(hexString)
	return resultString, nil
}
