package main

import (
	"fmt"
	"sync"
	"time"
)

// Função que simula uma tarefa concorrente
// Recebe um identificador (id) e um WaitGroup (wg) para sincronização
func tarefa(id int, wg *sync.WaitGroup) {
	defer func() {
		// O defer vai garantir que essa função será chamada no final da execução da função tarefa
		// wg.Done() avisa ao WaitGroup que esta goroutine terminou sua execução
		fmt.Printf("Tarefa %d: chamando wg.Done()\n", id)
		wg.Done()
	}()

	fmt.Printf("Tarefa %d começando\n", id) // Log para indicar o início da tarefa

	// Simula que a tarefa está "trabalhando" por 2 segundos, talvez um processamento ou I/O
	time.Sleep(2 * time.Second)

	// Log para indicar que a tarefa foi concluída
	fmt.Printf("Tarefa %d concluída\n", id)
}

func main() {
	var wg sync.WaitGroup // Cria uma instância do WaitGroup para gerenciar a sincronização das goroutines

	// Inicia três goroutines, cada uma executando a função tarefa
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Incrementa o contador do WaitGroup. Cada Add(1) indica que mais uma goroutine está em execução
		fmt.Printf("Main: adicionando Tarefa %d ao WaitGroup\n", i)
		go tarefa(i, &wg) // Inicia a função tarefa em uma nova goroutine
	}

	fmt.Println("Main: esperando todas as tarefas serem concluídas...")
	wg.Wait()                                        // O programa principal fica bloqueado aqui até que todas as goroutines tenham chamado wg.Done()
	fmt.Println("Todas as tarefas foram concluídas") // Mensagem final após todas as goroutines terminarem
}
