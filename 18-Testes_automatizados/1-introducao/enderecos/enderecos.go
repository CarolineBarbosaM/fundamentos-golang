package endereco

import "strings"

// TipoDeEndereco verifica se o endereco tem um tipo valido
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{
		"rua",
		"avenida",
		"estrada",
		"rodovia",
	}

	endEmLetraMaiuscula := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(endEmLetraMaiuscula, " ")[0]

	endTemTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			endTemTipoValido = true
		}
	}

	if endTemTipoValido {
		return strings.Title(primeiraPalavra)
	}

	return "Tipo Inválido"
}
