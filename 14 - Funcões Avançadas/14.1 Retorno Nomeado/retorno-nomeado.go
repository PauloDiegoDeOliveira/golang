package main

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	// Chamando a função e recebendo os valores de retorno.
	soma, subtracao := calculosMatematicos(10, 5)
	println("Soma:", soma)
	println("Subtração:", subtracao)
}
