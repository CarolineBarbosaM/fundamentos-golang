package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := variavel1 //a declaração é implicita, ela é do tipo que o valor é atribuido

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// declarando mais de uma variavel por vez
	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	)

	fmt.Println(variavel3, variavel4)

	// declarando duas variaveis
	// constantes não podem ser declaradas usando :=
	variavel5, variavel6 := "Variavel5", "Variavel6"

	fmt.Println(variavel5, variavel6)

	const constante1 string = "Contante1"
	fmt.Println(constante1)
}
