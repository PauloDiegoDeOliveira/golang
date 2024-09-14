package main

import (
	"fmt"
	"time"
)

// Função que simula uma tarefa demorada
func tarefa(nome string) {
	for i := 0; i < 3; i++ {
		fmt.Println(nome, "executando passo", i)
		time.Sleep(1 * time.Second) // Pausa por 1 segundo para simular tempo de execução
	}
}

func main() {
	// Iniciando duas goroutines concorrentes
	go tarefa("Tarefa 1") // Executa tarefa 1 em uma goroutine
	go tarefa("Tarefa 2") // Executa tarefa 2 em outra goroutine

	// A main goroutine precisa esperar as outras finalizarem
	time.Sleep(4 * time.Second)
	fmt.Println("Todas as tarefas concluídas")
}
