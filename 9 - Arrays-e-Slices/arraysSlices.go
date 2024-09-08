package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Definindo um array com tamanho fixo de 5 elementos do tipo int
	var array1 [5]int
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5

	// Exibe: [1 2 3 4 5]
	fmt.Println("Array 1:", array1)

	// Declarando e inicializando um array de tamanho 5
	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array 2:", array2)

	// Utilizando `...` para deixar o compilador determinar o tamanho do array
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array 3:", array3)

	// Declarando um slice sem elementos iniciais (slice vazio)
	var slice1 []int
	slice1 = append(slice1, 1, 2, 3, 4, 5)

	// Exibe: [1 2 3 4 5]
	fmt.Println("Slice 1:", slice1)

	// Declarando e inicializando um slice
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice 2:", slice2)

	// Criando um slice a partir de um array existente
	slice3 := array1[1:4] // slice contendo elementos do índice 1 ao 3 (exclusivo 4)
	fmt.Println("Slice 3 (a partir de array1):", slice3)

	// Utilizando make para criar um slice de tamanho 5 e capacidade 10
	slice4 := make([]int, 5, 10)
	// Inicializando os valores do slice
	for i := 0; i < len(slice4); i++ {
		slice4[i] = i + 1
	}
	fmt.Println("Slice 4 (com make):", slice4)

	// Exibindo o tipo do slice4 usando reflect
	fmt.Println(reflect.TypeOf(slice4)) // Isso mostra o tipo do slice4 (que será []int)

}
