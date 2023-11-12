package main

import "fmt"

func somar(num1 int8, num2 int8) int8 {
	return num1 + num2
}

// Funções podem ter mais de um retorno
func calculosMatematicos(num1, num2 int8) (int8, int8) {
	soma := num1 + num2
	subtracao := num1 - num2

	return soma, subtracao
}

func main() {
	soma := somar(30, 30)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Função F")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// Ignorando um resultado
	resultadoSoma2, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma2)
}
