package enderecos

import (
	"fmt"
	"strings"
)

// TipoEndereco identifica o tipo de um endereço (ex: Rua, Avenida, Alameda, etc.)
func TipoEndereco(endereco string) string {
	// Pega a primeira palavra do endereço
	primeiraPalavra := strings.Split(endereco, " ")[0]

	// Log para indicar qual foi a primeira palavra do endereço
	fmt.Println("A primeira palavra do endereço é:", primeiraPalavra)

	// Verifica o tipo de endereço
	switch primeiraPalavra {
	case "Rua", "Avenida", "Alameda":
		return primeiraPalavra
	default:
		return "Tipo Inválido"
	}
}
