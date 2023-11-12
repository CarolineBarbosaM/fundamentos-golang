package main

import (
	"fmt"
)

func main() {
	// Todos os dados dentro de um array devem ser do mesmo tipo
	// Senão especificar o tamanho do array ele vira slice
	var array [5]int
	array[0] = 1
	fmt.Println(array)

	array2 := [5]string{
		"Posição 0",
		"Posição 1",
		"Posição 2",
		"Posição 3",
		"Posição 4",
	}

	fmt.Println(array2)

	// O tamanho do array vai ser de acordo com a quantidade de valores
	// ATENÇÃO: Isso não deixa ele com tamanho dinamico
	array3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array3)

	// Slice é dinamico
	// Um slice aponta para um array (é uma fatia de um array)
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)

	slice1 = append(slice1, 18)
	fmt.Println(slice1)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Nova Posição"
	fmt.Println(slice2)

	//fmt.Println(reflect.TypeOf(array3))
	//fmt.Println(reflect.TypeOf(slice1))

	// Arrays Internos
	fmt.Println("------------------------------------")
	slice3 := make([]float32, 10, 15) // Slice, numero de posiçoes e capacidade
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
}
