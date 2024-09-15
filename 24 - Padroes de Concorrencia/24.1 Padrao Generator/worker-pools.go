package main

import (
	"fmt"
	"time"
)

// Worker (trabalhador) que recebe dados do canal de jobs e envia resultados para o canal de resultados
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs { // Recebe um job do canal de jobs
		fmt.Printf("[Worker %d] Iniciou o job %d\n", id, job)
		time.Sleep(time.Second) // Simula o tempo de processamento do job
		fmt.Printf("[Worker %d] Terminou o job %d\n", id, job)
		results <- job * 2 // Envia o resultado processado para o canal de resultados
	}
}

func main() {
	// Criando os canais de jobs (só recebe) e de resultados (só envia)
	// jobs <-chan int: só pode receber valores, usado pelos workers
	// results chan<- int: só pode enviar valores, usado pelos workers para enviar resultados
	jobs := make(chan int, 5)    // Canal de jobs com buffer de 5
	results := make(chan int, 5) // Canal de resultados com buffer de 5

	// Iniciando 3 workers (goroutines) que irão processar os jobs
	for w := 1; w <= 3; w++ {
		// Os workers só podem receber valores do canal de jobs e enviar para o canal de resultados
		go worker(w, jobs, results)
	}

	// Enviando 5 jobs para os workers
	for j := 1; j <= 5; j++ {
		fmt.Printf("[Main] Enviando job %d para o canal de jobs\n", j)
		jobs <- j // O canal de jobs só recebe valores, e os workers vão processar
	}
	close(jobs) // Fechamos o canal de jobs para indicar que não haverá mais trabalhos

	// Recebendo os resultados dos workers
	for r := 1; r <= 5; r++ {
		result := <-results // O canal de resultados só envia valores
		fmt.Printf("[Main] Resultado recebido: %d\n", result)
	}
}
