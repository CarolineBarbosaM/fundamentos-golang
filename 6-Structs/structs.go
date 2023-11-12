package main

import "fmt"

// Um struct tem como valor padrão zero {}
type usuario struct {
	nome     string
	idade    string
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo struct")

	var user usuario
	user.nome = "Maria"
	user.idade = "25"
	fmt.Println(user)

	end := endereco{"Rua Treis", 123}

	user2 := usuario{"João", "21", end}
	fmt.Println(user2)

	user3 := usuario{idade: "24"}
	fmt.Println(user3)
}
