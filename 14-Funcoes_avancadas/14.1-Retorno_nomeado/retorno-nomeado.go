package main

import "fmt"

// Retorno nomeado
// Não precisa declarar as variaveis, pois foram declaradas no retorno
func calculosMatematicos(n1, n2 int) (soma int, subratracao int) {
	soma = n1 + n2
	subratracao = n1 - n2

	return soma, subratracao
}

func main() {
	somar, subtrair := calculosMatematicos(5, 6)
	fmt.Println(somar, subtrair)
}
