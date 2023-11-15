// Teste de unidade

package endereco_test

import (
	endereco "introducao-a-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	enderecoEsperado string
}

// Se um teste der erro, tudo da erro
// Se rodou o teste uma vez (com sucesso), o go não roda novamente
// COMANDOS go test | go test ./... | go test -v | go test --cover
// Para gerar um relatorio dos testes use o comando: go test --coverprofile nomeDoArquivo.txt
// Para leitura do arquivo use o comando: go tool cover --func=nomeDoArquivo.txt
// Para saber o que não esta coberto, use: go tool cover --html=nomeDoArquivo.txt
func TestTipoDeEndereco(t *testing.T) {
	// Roda em paralelo com outros testes que tem a mesma funcao
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo Inválido"},
	}

	for _, cenarios := range cenariosDeTeste {
		tipoRecebido := endereco.TipoDeEndereco(cenarios.enderecoInserido)
		if tipoRecebido != cenarios.enderecoEsperado {
			t.Errorf("O tipo do erro é diferente do esperado!\nEsperava: %s | Recebeu: %s", cenarios.enderecoEsperado, tipoRecebido)
		}
	}
}
