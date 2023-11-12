package main

import "fmt"

func main() {
	// Todos os campos dentro do MAP tem que ter o mesmo tipo
	// Cuidado ao aninhar varios maps
	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "Da Silva",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Maria",
			"segundo":  "Silva",
		},
		"curso": {
			"tipo":   "Engenharia",
			"campus": "2",
		},
	}
	fmt.Println(usuario2["nome"])

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Leão",
	}

	fmt.Println(usuario2)
}
