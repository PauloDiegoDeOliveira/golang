package main

import "fmt"

// Função que modifica o valor de uma variável usando um ponteiro
func incrementarComPonteiro(x *int) {
	fmt.Println("Valor de x antes de incrementar:", *x)  // Mostra o valor apontado pelo ponteiro
	*x = *x + 1                                          // Modifica diretamente o valor na memória, usando o ponteiro
	fmt.Println("Valor de x depois de incrementar:", *x) // Mostra o valor atualizado
}

// Função que não usa ponteiros (cópia do valor)
func incrementarSemPonteiro(x int) {
	fmt.Println("Valor de x antes de incrementar (sem ponteiro):", x)
	x = x + 1 // Isso só modifica uma cópia local da variável
	fmt.Println("Valor de x depois de incrementar (sem ponteiro):", x)
}

func main() {
	// Exemplo usando ponteiro
	a := 10
	fmt.Println("Antes de chamar incrementarComPonteiro, valor de a:", a)
	incrementarComPonteiro(&a)                                             // Passando o endereço de 'a' para a função
	fmt.Println("Depois de chamar incrementarComPonteiro, valor de a:", a) // O valor de 'a' foi realmente modificado

	// Exemplo sem ponteiro
	b := 10
	fmt.Println("Antes de chamar incrementarSemPonteiro, valor de b:", b)
	incrementarSemPonteiro(b)                                              // Passando uma cópia de 'b' para a função
	fmt.Println("Depois de chamar incrementarSemPonteiro, valor de b:", b) // O valor de 'b' não foi modificado
}
