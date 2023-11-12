package main

import "fmt"

// Possui n parametros, ou seja, não precisa especificar a quantidade de parametros
// Os "..." tranforma em um slice. E se é um slice, vc pode interar
// Aceita também valores vazio
// Não da para usar mais de uma variatica na mesma funcao
func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	somar := soma(5, 6, 7, 8, 10, 15, 16, 17, 18)
	fmt.Println(somar)

	escrever("Olá Mundo!", 9, 11, 13, 16, 18, 20, 22)
}
