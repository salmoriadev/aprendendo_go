package main

import (
	"cryptopals/desafio01"
	"cryptopals/desafio02"
	"cryptopals/desafio03"
	"cryptopals/desafio04"
	"cryptopals/desafio05"

	//"cryptopals/desafio06"
	"cryptopals/desafio07"
	"cryptopals/desafio08"
)

func main() {
	resultado01, _ := desafio01.Desafio01()
	resultado02, _ := desafio02.Desafio02()
	resultado03, _ := desafio03.Desafio03()
	resultado04, _ := desafio04.Desafio04()
	resultado05, _ := desafio05.Desafio05()
	//resultado06, _ := desafio06.Desafio06()
	resultado07, _ := desafio07.Desafio07()
	resultado08, _ := desafio08.Desafio08()
	println("Desafio 01:", resultado01)
	println("Desafio 02:", resultado02)
	println("Desafio 03:", resultado03.Texto)
	println("Desafio 04:", resultado04.Texto)
	println("Desafio 05:", resultado05)
	//println("Desafio 06:", resultado06)
	println("Desafio 07:", resultado07)
	println("Desafio 08:", "Linha", resultado08.NumeroLinha, "Cifra:", resultado08.Cifra)
}
