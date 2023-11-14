package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		// O select é usado quando você possui uma função que demorar mais que a outra a finalizar
		// Usado somente em concorrência
		select {
		case msg1 := <-canal1:
			fmt.Println(msg1)

		case msg2 := <-canal2:
			fmt.Println(msg2)
		}
	}
}
