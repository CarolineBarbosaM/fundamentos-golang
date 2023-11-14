package main

import (
	"fmt"
	"time"
)

// Canais são um canal de comunicação podendo tanto enviar, quanto receber dados
// Cuidado com as deadlock, elas ocorrem quando não tem mais nenhum dado sendo enviado para o canal só que o canal ainda está esperando receber um dado
// O Go não consegue pegar um deadlock em compilação, somente em execução
func main() {
	canal := make(chan string)

	// Mandando o valor para o canal
	for {
		/*msg, aberto := <-canal
		// Verfica se o canal não estiver aberto, ele para a execução
		if !aberto {
			break
		}
		fmt.Println(msg)*/

		// Não precisa verificar pois o for vai até quando estiver abertos
		for msg := range canal {
			fmt.Println(msg)
		}

		fmt.Println("Fim do programa")
	}
}

func escrever(texto string, canal chan string) {
	for i := 0; 1 < 5; i++ {
		// Enviando pelo canal
		canal <- texto
		time.Sleep(time.Second)
	}

	// fecha o canal
	close(canal)
}
