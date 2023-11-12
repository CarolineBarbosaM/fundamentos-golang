package main

import "fmt"

func main() {
	// Funcoes sem nome
	func() {
		fmt.Println("Sem nome 1")
	}()

	func(texto string) {
		fmt.Println("Sem nome 2")
	}("Passando parametro")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando parametro")

	fmt.Println(retorno)
}
