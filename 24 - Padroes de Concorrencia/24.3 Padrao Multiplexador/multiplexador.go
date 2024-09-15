package main

import (
	"fmt"
	"time"
)

// Função que simula a geração de valores para um canal
// Esta função recebe um nome e um intervalo de tempo, e retorna um canal que vai gerar valores a cada intervalo
func generator(nome string, intervalo time.Duration) <-chan string {
	canal := make(chan string) // Criamos um canal que transmite strings

	// Iniciamos uma goroutine para enviar os valores para o canal de forma assíncrona
	go func() {
		// Laço que envia 5 valores para o canal
		for i := 1; i <= 5; i++ {
			// Envia uma mensagem formatada para o canal no formato "<nome>: valor <número>"
			canal <- fmt.Sprintf("%s: valor %d", nome, i)
			// Pausa a execução por um tempo para simular processamento
			time.Sleep(intervalo)
		}
		// Fecha o canal depois que todos os valores forem enviados
		close(canal)
	}()

	// Retorna o canal para que o consumidor possa receber os valores
	return canal
}

// Função que combina múltiplos canais em um único canal de saída (multiplexador)
// Recebe múltiplos canais de entrada como argumento (canais variáveis) e retorna um canal único
func multiplexar(canais ...<-chan string) <-chan string {
	saida := make(chan string) // Canal de saída que será retornado

	// Iniciamos uma goroutine para gerenciar a combinação dos canais
	go func() {
		// Variáveis para armazenar o estado de cada canal
		var aberto1, aberto2 bool
		var msg1, msg2 string

		// Loop infinito para ficar monitorando os canais
		for {
			// Usamos o select para aguardar dados de qualquer um dos canais
			select {
			// Caso o canal 1 envie uma mensagem, ele é tratado aqui
			case msg1, aberto1 = <-canais[0]:
				// Se o canal 1 ainda estiver aberto, enviamos a mensagem para o canal de saída
				if aberto1 {
					saida <- msg1
				}
			// Caso o canal 2 envie uma mensagem, ele é tratado aqui
			case msg2, aberto2 = <-canais[1]:
				// Se o canal 2 ainda estiver aberto, enviamos a mensagem para o canal de saída
				if aberto2 {
					saida <- msg2
				}
			}

			// Se ambos os canais forem fechados, fechamos também o canal de saída e terminamos a função
			if !aberto1 && !aberto2 {
				close(saida)
				return
			}
		}
	}()

	// Retorna o canal de saída para que os valores combinados possam ser consumidos
	return saida
}

func main() {
	// Criamos dois geradores de valores, que simulam fontes de dados com diferentes intervalos de tempo
	// O "Gerador 1" gera valores a cada 2 segundos
	canal1 := generator("Gerador 1", 2*time.Second)
	// O "Gerador 2" gera valores a cada 1 segundo
	canal2 := generator("Gerador 2", 1*time.Second)

	// Usamos a função multiplexar para combinar os dois canais em um único canal
	canalMultiplexado := multiplexar(canal1, canal2)

	// Consumimos os valores do canal multiplexado
	// O loop continua até que o canal multiplexado seja fechado
	for valor := range canalMultiplexado {
		// Imprimimos cada valor recebido do canal multiplexado
		fmt.Println("Recebido:", valor)
	}
}
