package main

import (
	"fmt"
)

// Função que recebe um número inteiro e imprime o dia da semana correspondente
func imprimirDiaDaSemana(dia int) {
	switch dia {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda-feira")
	case 3:
		fmt.Println("Terça-feira")
	case 4:
		fmt.Println("Quarta-feira")
	case 5:
		fmt.Println("Quinta-feira")
	case 6:
		fmt.Println("Sexta-feira")
	case 7:
		fmt.Println("Sábado")
	default:
		fmt.Println("Dia inválido")
	}
}

func main() {
	// Chamando a função com o valor 3 (Terça-feira)
	imprimirDiaDaSemana(3)

	// Você pode chamar a função com diferentes valores
	imprimirDiaDaSemana(1) // Domingo
	imprimirDiaDaSemana(5) // Quinta-feira
}
