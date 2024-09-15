package main

import (
	"fmt"
	"time"
)

func main() {
	// Criando dois canais, um sem buffer e outro com buffer
	canal1 := make(chan string)
	canal2 := make(chan string)

	// Goroutine para enviar dados no canal1
	go func() {
		time.Sleep(2 * time.Second) // Simulando algum processamento
		canal1 <- "Mensagem do canal1"
	}()

	// Goroutine para enviar dados no canal2
	go func() {
		time.Sleep(1 * time.Second) // Simulando um processamento mais rápido
		canal2 <- "Mensagem do canal2"
	}()

	// Usando select para aguardar mensagens de ambos os canais
	for i := 0; i < 2; i++ { // Espera por 2 mensagens, uma de cada canal
		select {
		case msg1 := <-canal1:
			// Quando o canal1 enviar a mensagem, essa linha será executada
			fmt.Println("Recebido:", msg1)
		case msg2 := <-canal2:
			// Quando o canal2 enviar a mensagem, essa linha será executada
			fmt.Println("Recebido:", msg2)
		}
	}
}
