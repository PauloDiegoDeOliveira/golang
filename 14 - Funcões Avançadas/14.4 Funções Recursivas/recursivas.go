package main

import "fmt"

// Função recursiva simples que calcula o enésimo número de Fibonacci
func fibonacciSimples(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciSimples(n-1) + fibonacciSimples(n-2)
}

// Função recursiva com memorização (usa um mapa para otimizar o cálculo)
var memo = make(map[int]int)

func fibonacciMemo(n int) int {
	if n <= 1 {
		return n
	}
	// Verifica se o valor já foi calculado
	if val, ok := memo[n]; ok {
		return val
	}
	// Calcula e armazena no mapa
	memo[n] = fibonacciMemo(n-1) + fibonacciMemo(n-2)
	return memo[n]
}

func main() {
	// Exemplo de uso da função recursiva simples
	fmt.Println("Fibonacci simples (primeiros 10 números):")
	for i := 0; i < 10; i++ {
		fmt.Printf("FibonacciSimples(%d) = %d\n", i, fibonacciSimples(i))
	}

	// Exemplo de uso da função com memorização
	numero := 40
	resultado := fibonacciMemo(numero)
	fmt.Printf("\nO %dº número de Fibonacci (com memorização) é: %d\n", numero, resultado)
}
