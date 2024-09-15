package tdd

import "fmt"

// Função Soma que recebe um slice de números e retorna a soma deles
func Soma(numeros []int) int {
	// Log para mostrar os números que estamos somando
	fmt.Printf("Somando os números: %v\n", numeros)

	// Variável para armazenar a soma
	total := 0

	// Itera sobre o slice de números e calcula a soma
	for _, numero := range numeros {
		total += numero
		// Log para mostrar a soma parcial
		fmt.Printf("Somando %d, total até agora: %d\n", numero, total)
	}

	// Retorna a soma total
	return total
}
