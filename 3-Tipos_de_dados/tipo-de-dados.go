package main

import (
	"errors"
	"fmt"
)

func main() {
	// Temos 4 tipos de numeros inteiros no Go
	// int8 int16 int32 int64
	// O numero é a quantida de bits que ele suporta: int8 até 8 bits
	// Caso use o int somente, ele irá usar com base na arquitetura do seu computador

	var numero int16 = 100
	fmt.Println(numero)

	var numero2 int32 = 1000000
	fmt.Println(numero2)

	var numero3 int = 1000000
	fmt.Println(numero3)

	// É o mesmo que int32, é só um alias
	var numero4 rune = 12334567
	fmt.Println(numero4)

	// É o mesmo que int8
	var numero5 byte = 123
	fmt.Println(numero5)

	// Temos 2 tipos de numeros real no Go
	// float32 float64
	// O numero é a quantida de bits que ele suporta: float32 até 32 bits

	var numeroReal1 float32 = 12.00
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1234568789.12
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// Tipo Texto
	// Não existe tipo "char" no go
	// Sempre usar aspas duplas ("")

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// O mais proximo do char que temos é saber a posição do caracter esta na tabela ascii
	char := 'C'
	fmt.Println(char)

	// Toda variavel em Go tem tipo inicial vazio/0
	var texto string
	fmt.Println(texto)

	// Tipo Logico
	var logico bool = true
	fmt.Println(logico)

	var logico2 bool
	fmt.Println(logico2) // false

	// Tipo erro
	// Para gerar erro no Go é necessário usar uma biblioteca
	var erro error = errors.New("Erro interno!")
	fmt.Println(erro)
}
