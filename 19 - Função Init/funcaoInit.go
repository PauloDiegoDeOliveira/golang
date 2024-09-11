package main

import "fmt"

// Variável global que será inicializada na função init
var contador int

// A função init é chamada automaticamente antes da função main
func init() {
	fmt.Println("Executando a função init")
	contador = 10 // Inicializando a variável global
	fmt.Println("Contador inicializado com valor:", contador)
}

func main() {
	fmt.Println("Executando a função main")
	// O valor de 'contador' já foi inicializado na função init
	fmt.Println("Valor do contador na função main:", contador)
}
