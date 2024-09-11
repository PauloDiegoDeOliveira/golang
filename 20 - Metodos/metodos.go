package main

import "fmt"

// Definindo uma struct chamada 'Retangulo'
type Retangulo struct {
	largura, altura float64
}

// Método para calcular a área do Retangulo (receiver por valor)
// Este método não altera o valor do retângulo, pois o receiver é passado por valor.
func (r Retangulo) area() float64 {
	return r.largura * r.altura
}

// Método que modifica o valor do Retangulo (receiver por ponteiro)
// Como estamos usando um ponteiro, este método pode alterar diretamente o valor do objeto original.
func (r *Retangulo) dobrarTamanho() {
	r.largura *= 2
	r.altura *= 2
}

// Função que recebe um ponteiro para um número inteiro e o modifica diretamente
func incrementarNumero(num *int) {
	fmt.Println("Valor de num antes do incremento:", *num)  // Mostra o valor atual
	*num += 1                                               // Incrementa o valor apontado pelo ponteiro
	fmt.Println("Valor de num depois do incremento:", *num) // Mostra o valor atualizado
}

func main() {
	// Exemplo com a struct 'Retangulo'
	r := Retangulo{largura: 10, altura: 5}

	// Chamando o método 'area' para calcular a área
	fmt.Println("Área do retângulo:", r.area()) // Saída: 50

	// Chamando o método 'dobrarTamanho' que altera o valor da struct
	r.dobrarTamanho()
	fmt.Println("Novas dimensões - Largura:", r.largura, "Altura:", r.altura) // Saída: Largura: 20, Altura: 10

	// Calculando a nova área após dobrar o tamanho
	fmt.Println("Nova área do retângulo:", r.area()) // Saída: 200

	// Exemplo com ponteiro para um inteiro
	numero := 10
	fmt.Println("Valor inicial de numero:", numero) // Exibe o valor inicial de 'numero'

	// Chamando a função 'incrementarNumero' passando o ponteiro de 'numero'
	incrementarNumero(&numero) // Passa o endereço de 'numero' para que seja modificado diretamente

	// O valor de 'numero' foi modificado pela função, já que passamos um ponteiro
	fmt.Println("Valor final de numero após incremento:", numero) // Saída: 11
}
