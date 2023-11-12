package main

import "fmt"

func main() {
	// Operadores Aritmeticos
	// Não é possivel fazer calculos com tipos diferentes: int16 + int32
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 2
	multiplicacao := 10 * 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// Atribuições
	var atribuicao string = "Atribuição"
	atribuicao2 := "Atribuição 2"

	fmt.Println(atribuicao, atribuicao2)

	// Operadores Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// Operadores Logicos
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Operadores Unarios
	// Só interagem com uma variavel por vez
	numero := 10

	numero++
	numero += 15
	numero--
	numero -= 20

	numero *= 2
	numero /= 20
	numero %= 2

	fmt.Println(numero)

	// Operador Tenario
	// Não existe operador ternario em Go, pois ele tem como primicia é que só tem uma meneira de fazer cada coisa
	// Use o if e else para esse cenários
}
