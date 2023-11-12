package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Implementando interação i")
		time.Sleep(time.Second)
	}

	for j := 0; j < 10; j++ {
		fmt.Println("Implementando interação j", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	// Retorna o codigo da letra na tabela ascci
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Maria",
		"sobrenome": "Oliveira",
	}

	fmt.Println(usuario)

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// Não da para usar em struct

	//Loop infinito
	//for true{
	//fmt.Println("Execução infinita")
	//time.Sleep(time.Second)
	//}
}
