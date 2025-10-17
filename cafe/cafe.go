package main

import (
	"fmt"
)

type Cafe struct {
	ProdutoBase
	Moagem string
}

func (c Cafe) VerAtributos() string {
	return "Café: " + c.Nome + ", Preço de Venda: R$ " + fmt.Sprintf("%.2f", c.PrecoVenda) + ", Lucro: R$ " + fmt.Sprintf("%.2f", c.CalcularLucro()) + ", Moagem: " + c.Moagem
}
