package main

import (
	"fmt"              // Importa o pacote fmt para imprimir na saída
	"testes/enderecos" // Certifique-se de que o caminho de importação corresponde ao seu módulo
)

func main() {
	// Exemplo de uso da função do pacote enderecos
	endereco := "Alameda dos Anjos 456" // Definindo uma string que representa um endereço

	// Log para informar que estamos começando a identificar o tipo de endereço
	fmt.Println("Iniciando identificação do tipo de endereço:", endereco)

	// Chama a função TipoEndereco do pacote enderecos para obter o tipo do endereço
	tipo := enderecos.TipoEndereco(endereco)

	// Log para informar qual foi o tipo de endereço retornado pela função
	fmt.Println("O tipo de endereço identificado foi:", tipo)

	// Imprime o tipo do endereço final para o usuário
	fmt.Println("O tipo do endereço é:", tipo)
}
