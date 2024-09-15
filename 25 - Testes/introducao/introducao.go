package main

import (
	"fmt"
	"testes/enderecos" // Certifique-se de que o caminho de importação corresponde ao módulo
)

func main() {
	// Exemplo de uso da função
	endereco := "Alameda dos Anjos 456"
	tipo := enderecos.TipoEndereco(endereco) // Chama a função TipoEndereco do pacote enderecos
	fmt.Println("O tipo do endereço é:", tipo)
}
