package main

import "fmt"

func main() {
	// Em um canal com buffer, você especifica a capacidade do seu canal
	canal := make(chan string, 2)
	canal <- "Olá Mundo\n"
	canal <- "Programando em Go\n"
	//canal <- "Novo valor" // Aqui já vai dar deadlock

	msg := <-canal
	msg2 := <-canal
	fmt.Println(msg, msg2)
}
