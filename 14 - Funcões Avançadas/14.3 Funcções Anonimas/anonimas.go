package main

import (
	"fmt"
)

func main() {
	// Define uma função anônima que recebe um texto, realiza a subtração e exibe o resultado formatado
	func(texto string) {
		a := 10
		b := 5
		subtracao := a - b
		// Usa Sprintf para formatar a string com o texto e o resultado da subtração
		resultadoFormatado := fmt.Sprintf("%s %d", texto, subtracao)
		fmt.Println(resultadoFormatado) // Exibe a string formatada com fmt.Println
	}("Subtração formatada é:")
}
