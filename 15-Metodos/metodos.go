package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// Trata ações
func (u usuario) salvar() {
	fmt.Printf("Salvando usuario %s no banco de dados.\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) aniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Neuza", 86}
	usuario1.salvar()

	usuario2 := usuario{"Mario", 15}
	usuario2.salvar()

	maiorDeIdade := usuario1.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.aniversario()
	fmt.Println(usuario2.idade)
}
