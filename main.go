package main

import (
	"app-endereco-com-teste/endereco"
	"fmt"
)

func main() {
	tipoEndereco := endereco.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)

}