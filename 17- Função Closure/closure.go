package main

import "fmt"

// Função closure que gera um contador
func contador() func() int {
	// 'i' está definido dentro do escopo de 'contador' e será "capturado" pelo closure
	i := 0
	return func() int {
		// Toda vez que a função retornada for chamada, 'i' será incrementado
		i++
		fmt.Println("Valor de i dentro do closure:", i)
		return i
	}
}

func main() {
	// Primeiro closure: 'incremento' é um closure que manipula seu próprio 'i'
	incremento := contador()

	fmt.Println("Primeira chamada do primeiro closure:")
	fmt.Println(incremento()) // i começa em 0 e é incrementado para 1

	fmt.Println("Segunda chamada do primeiro closure:")
	fmt.Println(incremento()) // i agora é 1 e é incrementado para 2

	fmt.Println("Terceira chamada do primeiro closure:")
	fmt.Println(incremento()) // i agora é 2 e é incrementado para 3

	// Segundo closure: 'outroIncremento' tem seu próprio 'i', independente do primeiro closure
	outroIncremento := contador()

	fmt.Println("Primeira chamada do segundo closure:")
	fmt.Println(outroIncremento()) // i começa em 0 novamente (novo closure)

	fmt.Println("Segunda chamada do segundo closure:")
	fmt.Println(outroIncremento()) // i agora é 1 e é incrementado para 2

	// O primeiro closure ainda "lembra" do seu valor de 'i'
	fmt.Println("Quarta chamada do primeiro closure:")
	fmt.Println(incremento()) // i ainda é 3 e é incrementado para 4
}
