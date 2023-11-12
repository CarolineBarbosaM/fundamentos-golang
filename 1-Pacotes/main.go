package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo no arquivo main!")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat(("123")) // para remover o pacote use: go mod tidy
	fmt.Println(erro)
}
