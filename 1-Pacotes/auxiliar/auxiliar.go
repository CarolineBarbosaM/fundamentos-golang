package auxiliar

import "fmt"

// Escrever: registra nome na tela
func Escrever() {
	/*
	* Se a função tem letra minuscula, ela é privada.
	* Se a função tem letra maiuscula, ela é publica.
	* Toda função publica (pode exportar) por boa pratica, deve ter um comentário em cima dizendo o que ela faz
	 */
	fmt.Println("Escrevendo do arquivo do pacote auxiliar.")
	escrever2()

}
