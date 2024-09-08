package main

import (
	"fmt"
	"net/http"
	"time"
)

// Estrutura que representa uma tarefa
type Task struct {
	ID          int
	Description string
}

// Função que processa uma tarefa (simula um processamento demorado)
func processTask(task Task) {
	fmt.Printf("Iniciando processamento da tarefa #%d: %s\n", task.ID, task.Description)
	time.Sleep(2 * time.Second) // Simula o tempo de processamento
	fmt.Printf("Tarefa #%d concluída\n", task.ID)
}

// Função handler que será chamada para processar requisições HTTP
func taskHandler(w http.ResponseWriter, r *http.Request) {
	// Cria uma tarefa para processar
	task := Task{ID: 1, Description: "Enviar relatório"}
	go processTask(task) // Executa o processamento em uma goroutine (concorrente)

	// Responde ao cliente imediatamente
	fmt.Fprintf(w, "Tarefa #%d está sendo processada\n", task.ID)
}

// 1. Função Simples (com parâmetros e retorno)
func somar(n1 int8, n2 int8) int8 { // Função que recebe dois parâmetros e retorna a soma
	return n1 + n2
}

// 2. Função com múltiplos retornos
func calculosMatematicos(n1, n2 int8) (int8, int8) { // Função com dois retornos
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

// 3. Função sem parâmetros e sem retorno
func exibirMensagem() { // Função simples sem parâmetros nem retorno
	fmt.Println("Mensagem simples")
}

// 10. Métodos (funções associadas a tipos)
type Retangulo struct { // Definindo um tipo Retângulo
	largura, altura int
}

func (r Retangulo) area() int { // Método associado ao tipo Retangulo
	return r.largura * r.altura
}

func main() {
	// Função anônima atribuída a uma variável
	var saudacao = func(nome string) { // Função anônima recebendo um parâmetro
		fmt.Println("Olá", nome)
	}
	saudacao("Paulo Diego") // Chamando a função anônima

	// Função anônima executada diretamente
	func() {
		fmt.Println("Função anônima executada diretamente")
	}()

	// 5. Função com retorno nomeado
	dividir := func(n1, n2 float64) (resultado float64, erro error) {
		if n2 == 0 {
			erro = fmt.Errorf("não é possível dividir por zero")
			return
		}
		resultado = n1 / n2
		return // Aqui retorna 'resultado' e 'erro'
	}
	resultado, erro := dividir(10, 2)
	if erro != nil {
		fmt.Println(erro)
	} else {
		fmt.Println("Resultado da divisão:", resultado)
	}

	// 6. Funções Variádicas (quantidade variável de argumentos)
	somarTudo := func(numeros ...int) int { // Função que aceita um número variável de argumentos
		total := 0
		for _, numero := range numeros {
			total += numero
		}
		return total
	}

	fmt.Println(somarTudo(1, 2, 3, 4, 5)) // Passando múltiplos argumentos

	// 7. Funções como parâmetros
	executarOperacao := func(n1 int, n2 int, operacao func(int, int) int) int { // Função que recebe outra função como parâmetro
		return operacao(n1, n2)
	}

	somar := func(n1, n2 int) int { // Função normal de soma
		return n1 + n2
	}

	resultadoOperacao := executarOperacao(10, 20, somar) // Passando a função 'somar' como parâmetro
	fmt.Println("Resultado da operação:", resultadoOperacao)

	// 8. Funções como retorno
	obterOperacao := func(operacao string) func(int, int) int { // Função que retorna outra função
		if operacao == "somar" {
			return func(a, b int) int {
				return a + b
			}
		} else if operacao == "subtrair" {
			return func(a, b int) int {
				return a - b
			}
		}
		return nil
	}

	op := obterOperacao("somar") // Função retornada que realiza soma
	fmt.Println(op(10, 20))      // Resultado: 30

	// 9. Closures (funções que capturam variáveis externas)
	contador := func() func() int { // Função que retorna outra função (closure)
		cont := 0
		return func() int {
			cont++
			return cont
		}
	}

	contar := contador()  // Função contador
	fmt.Println(contar()) // 1
	fmt.Println(contar()) // 2

	// Usando o método area da estrutura Retangulo
	ret := Retangulo{10, 20}         // Criando uma instância de Retangulo
	fmt.Println("Área:", ret.area()) // Chamando o método 'area'

	// Define o handler para a rota "/task"
	http.HandleFunc("/task", taskHandler)

	// Inicia o servidor HTTP na porta 8080
	fmt.Println("Servidor iniciado na porta 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
