package main

import "fmt"

// Muito cuidado ao usar
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("Texto")
	generica(1)
	generica(true)

	// NÃO FAÇA ISSOOO
	mapa := map[interface{}]interface{}{
		1:            "Texto",
		true:         float64(54.7),
		"OutroTexto": false,
	}

	fmt.Println(mapa)
}
