package main

func main() {
	// Definindo uma variável chamada "numero" e atribuindo o valor 10 a ela
	numero := 10

	// Primeira estrutura de controle "if"
	if numero > 15 {
		// Se o valor de "numero" for maior que 15, imprime esta mensagem
		println("Maior que 15")
	} else {
		// Caso contrário, se o valor de "numero" for menor ou igual a 15, imprime esta mensagem
		println("Menor que 15")
	}

	// Segunda estrutura de controle "if" com declaração dentro da condição
	if outroNumero := numero; outroNumero > 0 {
		// Aqui, criamos uma nova variável "outroNumero", com o mesmo valor de "numero"
		// Se "outroNumero" for maior que 0, imprime esta mensagem
		println("Número é maior que zero")
	} else {
		// Caso contrário, se "outroNumero" for menor ou igual a 0, imprime esta mensagem
		println("Número é menor que zero")
	}
}
