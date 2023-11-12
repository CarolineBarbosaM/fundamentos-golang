package main

import "fmt"

// Deve ter um condição para finalizar, senão acontece o estouro de pilha
// Estouro de pilha é quando a pilha fica muito grande e estoura
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

// São funções que chamam elas mesmo
// Não são recomendadas quando houver muitas chamadas
func main() {
	posicao := uint(10)

	for i := uint(0); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
