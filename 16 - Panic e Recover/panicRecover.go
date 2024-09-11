package main

import "fmt"

// Função que verifica se o aluno está aprovado
func alunoEstaAprovado(n1, n2 float32) bool {
	defer func() { // O defer garante que o código dentro dessa função anônima será executado no final, mesmo após o panic
		if r := recover(); r != nil { // O recover captura o valor passado para o panic e impede que o programa seja interrompido
			fmt.Println("Recuperado do erro:", r)
		}
	}()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é exatamente 6") // O panic interrompe o fluxo normal do programa
}

func main() {
	// Primeira chamada da função sem problemas
	fmt.Println(alunoEstaAprovado(7, 8)) // true

	// Segunda chamada da função que causará o panic
	fmt.Println(alunoEstaAprovado(6, 6)) // panic, mas será recuperado pelo recover

	fmt.Println("Programa continua normalmente após o panic ser recuperado") // Essa linha será executada graças ao recover
}
