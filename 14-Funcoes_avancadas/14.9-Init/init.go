package main

import "fmt"

var n int

// Função que executada antes da função main
// A ordem não altera a chamada
// Só pode ter UMA por arquivo
func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada.")
	fmt.Println(n)
}
