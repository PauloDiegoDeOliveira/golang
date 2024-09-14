package main

import (
	"fmt"
)

// Função que envia múltiplos valores para o canal
func enviarValores(canal chan int, metodo string) {
	for i := 1; i <= 5; i++ { // Enviando 5 valores através do canal
		fmt.Printf("[%s] Enviando valor para o canal: %d\n", metodo, i)
		canal <- i
	}
	fmt.Printf("[%s] Todos os valores enviados, fechando o canal.\n", metodo)
	close(canal) // Fechando o canal após o envio dos valores
}

// Função que recebe os valores usando for com if/else
func receberComFor(canal chan int) {
	fmt.Println("[For com if/else] Iniciando recebimento de valores.")
	// Recebe 5 valores do canal
	for i := 1; i <= 5; i++ {
		valorRecebido := <-canal // Recebe o valor do canal
		fmt.Printf("[For com if/else] Mensagem %d: Valor recebido do canal: %d\n", i, valorRecebido)
	}

	// Verifica se o canal está fechado
	_, aberto := <-canal
	if !aberto {
		fmt.Println("[For com if/else] Canal fechado após receber todas as mensagens.")
	} else {
		fmt.Println("[For com if/else] O canal ainda está aberto.")
	}
}

// Função que recebe os valores usando range
func receberComRange(canal chan int) {
	fmt.Println("[Range] Iniciando recebimento de valores.")
	// Usando range para receber os valores do canal até que ele seja fechado
	for valorRecebido := range canal {
		fmt.Printf("[Range] Valor recebido do canal: %d\n", valorRecebido)
	}

	fmt.Println("[Range] Canal fechado após receber todas as mensagens.")
}

func main() {
	// Primeiro exemplo: Usando for com if/else
	fmt.Println("Exemplo 1: Receber valores com for e if/else")
	canal1 := make(chan int)

	// Inicia a função enviarValores em uma goroutine
	go enviarValores(canal1, "For com if/else")

	// Recebe os valores usando o for com if/else
	receberComFor(canal1)

	// Segundo exemplo: Usando range
	fmt.Println("\nExemplo 2: Receber valores com range")
	canal2 := make(chan int)

	// Inicia a função enviarValores em uma goroutine
	go enviarValores(canal2, "Range")

	// Recebe os valores usando o range
	receberComRange(canal2)
}
