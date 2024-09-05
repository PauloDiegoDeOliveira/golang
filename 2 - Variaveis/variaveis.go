package main

import (
	"fmt"
)

func main() {

	var variavel1 string = "Variável 1" // Declaração de variáveis com tipagem explícita

	variavel2 := "Variável 2" // Declaração de variáveis com tipagem implícita

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// Declaração de múltiplas variáveis
	var (
		variavel3 string = "Variável 1"
		variavel4 string = "Variável 2"
	)

	fmt.Println(variavel3, variavel4)

	// Declaração de múltiplas variáveis com tipagem implícita
	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1" // Declaração de constantes
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5 // Troca de valores entre variáveis
	fmt.Println(variavel5, variavel6)
}
