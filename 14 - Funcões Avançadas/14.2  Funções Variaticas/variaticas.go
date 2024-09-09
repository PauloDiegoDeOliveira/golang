package main

import "fmt"

// Função variádica que aceita uma string e múltiplos inteiros
func processaNumeros(mensagem string, numeros ...int) int {
	total := 0 // Inicializa a variável 'total' com valor 0

	// Loop que percorre cada valor passado para a função
	for i, v := range numeros {
		// Exibe a mensagem personalizada com o valor atual e o índice
		fmt.Printf("%s no índice %d: %d\n", mensagem, i, v)
		total += v // Soma o valor atual ao total
	}

	return total // Retorna o valor final da soma
}

func main() {
	// Chama a função variádica com uma string e diferentes quantidades de inteiros
	somaTotal := processaNumeros("Processando número", 1, 2, 3, 4, 5)
	fmt.Println("Soma total é:", somaTotal) // Deve imprimir a soma total

	// Outra chamada com valores diferentes
	somaTotal = processaNumeros("Processando número", 10, 20)
	fmt.Println("Soma total é:", somaTotal) // Deve imprimir a soma total

	// Chamada sem números
	somaTotal = processaNumeros("Processando número") // Nenhum número passado
	fmt.Println("Soma total é:", somaTotal)           // Deve imprimir 0
}
