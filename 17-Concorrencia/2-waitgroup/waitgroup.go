package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// O Wait Group é umas das formas de sicronizar goroutines
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Olá Mundo!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em Go")
		waitGroup.Done()
	}()

	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 0; 1 < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
