package main

import (
	"fmt"
	"time"
)

func main() {
	// Criando um canal com buffer de tamanho 3
	// Isso significa que o canal pode armazenar até 3 valores antes de bloquear a goroutine que está enviando.
	canal := make(chan int, 3)

	// Função anônima para enviar valores para o canal
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Enviando valor %d para o canal...\n", i)
			canal <- i // Envia o valor para o canal. Se o canal estiver cheio (mais que 3 valores), ele vai bloquear até que haja espaço.
			fmt.Printf("Valor %d enviado.\n", i)
			time.Sleep(1 * time.Second) // Pausa para simular tempo de processamento
		}
		close(canal) // Fecha o canal após enviar todos os valores. Isso é importante para sinalizar que não haverá mais dados.
	}()

	// Recebe valores do canal
	// Como estamos recebendo os valores mais lentamente, podemos ver o comportamento do buffer.
	for valor := range canal {
		fmt.Printf("Valor recebido: %d\n", valor)
		time.Sleep(2 * time.Second) // Simulando processamento mais lento ao receber os dados
	}
}
