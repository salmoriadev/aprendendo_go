package main

import (
	"fmt"
)

type MaquinaCafe struct {
	ProdutoBase
	Capacidade int
}

func (m MaquinaCafe) VerAtributos() string {
	return "Máquina de Café: " + m.Nome + ", Preço de Venda: R$ " + fmt.Sprintf("%.2f", m.PrecoVenda) + ", Lucro: R$ " + fmt.Sprintf("%.2f", m.CalcularLucro()) + ", Capacidade: " + fmt.Sprintf("%dML", m.Capacidade)
}
