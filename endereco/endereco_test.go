package endereco_test

import (
	. "app-endereco-com-teste/endereco"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipodeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste {
		{"Avenida Juarez Frei Barbosa", "Avenida"},
		{"Rua Vigilio Fonseca", "Rua"},
		{"Sitio Floresta", "Sitio"},
		{"Rodovia do Pimenta", "Rodovia"},
		{"Praça Municipal Frei Luis", "Tipo inválido!"},
		{"Estrada Luis de Sá", "Estrada"},
		{"RUA MARIA DE MUNIZ", "Rua"},
		{"avenida rebouças", "Avenida"},
		{"", "Tipo inválido!"},
		{"Estrada PAULO 2", "Estrada"},
		{"SitIO Paulo MACHADO 4", "Sitio"},
		{"663", "Tipo inválido!"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != TipoDeEndereco(cenario.retornoEsperado) {
			t.Errorf("O tipo recebido é diferente do esperado - %s ", cenario.retornoEsperado)
		}
	}
}