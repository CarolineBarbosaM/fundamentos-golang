package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunosRepovados(n1, n2 float32) bool {
	defer fmt.Println("Media calculada. O resultado será retornado!")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

// Defer adia a execução de um determinado pedaço de código
func main() {
	fmt.Println(alunosRepovados(8, 3))
}
