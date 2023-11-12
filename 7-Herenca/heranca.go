package main

import "fmt"

type pessoas struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudantes struct {
	pessoas
	curso     string
	faculdade string
}

// Em Go não existe herança do jeito tradicional
// Esse é o formato mais proximo de herança que temos
// Você consegue criar em GO structs não usadas, pois structs são um tipo
func main() {
	fmt.Println("Herança")

	person1 := pessoas{"Maria", "Da Silva", 24, 157}
	fmt.Println(person1)

	student := estudantes{person1, "Arquitetura", "PUC"}
	fmt.Println(student)
	fmt.Println(student.nome)
}
