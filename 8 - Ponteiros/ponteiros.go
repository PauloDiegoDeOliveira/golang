package main

import (
	"fmt"
)

func main() {

	variavel1 := 10
	variavel2 := variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// Ponteiro é uma referência de memória
	var variavel3 int
	var ponteiro *int // * é um ponteiro

	variavel3 = 100
	ponteiro = &variavel3 // & referenciando o ponteiro

	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) // * desreferenciando o ponteiro

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)
}
