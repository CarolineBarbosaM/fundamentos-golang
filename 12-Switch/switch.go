package main

import "fmt"

func diasDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"

	case 2:
		return "Segunda-Feira"

	case 3:
		return "Terça-Feira"

	case 4:
		return "Quarta-Feira"

	case 5:
		return "Quinta-Feira"

	case 6:
		return "Sexta-Feira"

	case 7:
		return "Sabado"

	default:
		return "Numero Inválido"
	}
}

func mesesDoAno(numero int) string {

	var mesDoAno string

	switch {
	case numero == 1:
		mesDoAno = "Janeiro"
		fallthrough // É usado somente em casos muito especificos | Joga o codigo para a proxima condição

	case numero == 2:
		mesDoAno = "Fevereiro"

	case numero == 3:
		mesDoAno = "Março"

	case numero == 4:
		mesDoAno = "Abril"

	case numero == 5:
		mesDoAno = "Maio"

	case numero == 6:
		mesDoAno = "Junho"

	case numero == 7:
		mesDoAno = "Julho"

	case numero == 8:
		mesDoAno = "Agosto"

	case numero == 9:
		mesDoAno = "Setembro"

	case numero == 10:
		mesDoAno = "Outubro"

	case numero == 11:
		mesDoAno = "Novembro"

	case numero == 12:
		mesDoAno = "Dezembro"

	default:
		mesDoAno = "Numero Inválido"
	}

	return mesDoAno
}

func main() {
	dia := diasDaSemana(3)
	fmt.Println(dia)

	mes := mesesDoAno(9)
	fmt.Println(mes)
}
