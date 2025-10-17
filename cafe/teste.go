package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite o nome do café: ")
	nomeCafe, _ := reader.ReadString('\n')
	nomeCafe = strings.TrimSpace(nomeCafe)

	fmt.Print("Digite o preço de venda do café: ")
	var precoVendaCafe float64
	fmt.Scanf("%f\n", &precoVendaCafe)

	fmt.Print("Digite o preço de compra do café: ")
	var precoCompraCafe float64
	fmt.Scanf("%f\n", &precoCompraCafe)

	cafe := Cafe{
		ProdutoBase: ProdutoBase{
			Nome:        nomeCafe,
			PrecoVenda:  precoVendaCafe,
			PrecoCompra: precoCompraCafe,
		},
		Moagem: "Grão",
	}

	fmt.Print("\nAgora, digite o modelo da máquina de café: ")
	modeloMaquina, _ := reader.ReadString('\n')
	modeloMaquina = strings.TrimSpace(modeloMaquina)

	fmt.Print("Digite a capacidade da máquina de café (em ML): ")
	var capacidadeMaquina int
	fmt.Scanf("%d\n", &capacidadeMaquina)

	maquina := MaquinaCafe{
		Capacidade: capacidadeMaquina,
		ProdutoBase: ProdutoBase{
			Nome:        modeloMaquina,
			PrecoVenda:  300.0,
			PrecoCompra: 250.0,
		},
	}

	fmt.Println("\nDetalhes do Café:")
	fmt.Println(cafe.VerAtributos())

	fmt.Println("\nDetalhes da Máquina de Café:")
	fmt.Println(maquina.VerAtributos())
}
