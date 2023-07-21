package endereco

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "sitio", "rodovia", "estrada"}
	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTipoValido = true
		}
	}
	
	if enderecoTipoValido {
		caser := cases.Title(language.BrazilianPortuguese)
		return caser.String(primeiraPalavraDoEndereco)
	}

	return "Tipo inválido!"
}