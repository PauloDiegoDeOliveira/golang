package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := 10000000000 // int
	fmt.Println(numero)

	var numero2 uint32 = 100000 // uint32
	fmt.Println(numero2)

	// alias
	var numero3 rune = 12456 // int32 = rune - numeros que representam caracteres
	fmt.Println(numero3)

	var numero4 byte = 123 // byte = uint8
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45 // float32
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123 // float64
	fmt.Println(numeroReal2)

	numeroReal3 := 123.45 // float64
	fmt.Println(numeroReal3)

	var str string = "Texto" // string
	fmt.Println(str)

	str2 := "Texto2" // string
	fmt.Println(str2)

	char := 'B' // char
	fmt.Println(char)

	char2 := 'A' // char
	fmt.Println(char2)

	var texto string
	fmt.Println(texto) // string vazia = ""

	var Texto2 int16
	fmt.Println(Texto2) // int16 vazio = 0

	numero5 := 5
	fmt.Println(numero5) // int

	var booleano1 bool = true // bool
	fmt.Println(booleano1)

	var booleano2 bool // bool
	fmt.Println(booleano2)

	var erro error
	fmt.Println(erro) // error vazio = nil

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)

}
