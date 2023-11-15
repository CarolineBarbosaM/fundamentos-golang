package main

import (
	"fmt"
	endereco "introducao-a-testes/enderecos"
)

func main() {
	tipo_de_endereco := endereco.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipo_de_endereco)
}
