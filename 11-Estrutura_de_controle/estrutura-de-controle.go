package main

import "fmt"

func main() {
	numero := -5

	// Condiçoes não funcionam sem chave
	if numero > 15 {
		fmt.Println("Maior do que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// if init
	// a variavel declarada só existe dentro do if
	if outronumero := numero; outronumero > 0 {
		fmt.Println("Numero é maior que 0")
	} else {
		fmt.Println("Numero é menor que 0")
	}

	// aninhado
	if outronumero := numero; outronumero > 0 {
		fmt.Println("Numero é maior que 0")
	} else if numero < -10 {
		fmt.Println("Numero é menor que -10")
	} else {
		fmt.Println("Numero é menor que 0")
	}
}
