package enderecos

import (
	"strings"
)

// Função que retorna o tipo de endereço com base nas palavras
func TipoEndereco(endereco string) string {
	tiposValidos := map[string]string{
		"Rua":      "Residencial",
		"Avenida":  "Comercial",
		"Rodovia":  "Rodovia",
		"Via":      "Via",
		"Alameda":  "Residencial",
		"Estrada":  "Estrada",
		"Travessa": "Travessa",
		"Quadra":   "Residencial",
	}

	palavras := strings.Split(endereco, " ")

	for _, palavra := range palavras {
		if tipo, existe := tiposValidos[palavra]; existe {
			return tipo
		}
	}

	return "Tipo desconhecido"
}
