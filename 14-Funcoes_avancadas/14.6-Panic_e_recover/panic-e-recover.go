package main

import "fmt"

// Recupera a execução de panic
func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada!")
	}
}

// O Panic para de executar e entra em panico
func alunosRepovados(n1, n2 float32) bool {
	defer recuperarExecucao()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	// Não é um erro
	panic("A MÉDIA É EXATAMENTE 6")
}

func main() {
	fmt.Println(alunosRepovados(6, 6))
	fmt.Println("Pós execução")
}
