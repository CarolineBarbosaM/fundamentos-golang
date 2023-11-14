package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá Mundo"), escrever("Programando em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// Coloca dois canais em um canal só
func multiplexar(canalDeEntrada1, canaldeEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case msg := <-canalDeEntrada1:
				canalDeSaida <- msg
			case msg := <-canaldeEntrada2:
				canalDeSaida <- msg
			}
		}
	}()

	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
