package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARALELISMO
	// PARALELISMO: Você duas ou mais tarefas sendo executada ao mesmo tempo
	// CONCORRÊNCIA: Não executa ao mesmo tempo necesseriamente as tarefas, uma tarefa só não espera outra acabar para executar.
	// GOROUTINE é uma função ou metodo que não esperar a finalização de uma tarefa para executar a proxima
	go escrever("Olá Mundo!")
	escrever("Programando em Go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
