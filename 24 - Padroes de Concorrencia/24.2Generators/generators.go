package main

import (
	"fmt"
)

// Função generator que gera uma sequência de números inteiros
func generator() <-chan int {
	// Criando um canal sem buffer
	canal := make(chan int)

	// Iniciando uma goroutine que envia os valores para o canal
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("[Generator] Enviando valor: %d\n", i)
			canal <- i // Enviando o valor para o canal
		}
		close(canal) // Fecha o canal após enviar todos os valores
	}()

	// Retorna o canal para que o consumidor possa receber os valores
	return canal
}

func main() {
	// Chamando a função generator que retorna um canal
	canal := generator()

	// Consumindo os valores gerados pelo generator
	for valor := range canal {
		fmt.Printf("[Main] Valor recebido: %d\n", valor)
	}
}
