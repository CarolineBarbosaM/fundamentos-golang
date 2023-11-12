package main

import "fmt"

// Ponteiros são uma referência de memoria
// Quando você cria um ponteiro e atribui valor, você não atribui o valor, mas o endereço de memoria
func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++

	fmt.Println(variavel1, variavel2)

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3 // endereço de memoria da variavel 3

	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) // Desferenciação: remover a referência
}
