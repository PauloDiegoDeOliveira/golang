package main

func main() {
	// Exemplo 1: Loop "for" com inicialização, condição e incremento.
	// Imprime os valores de 0 a 9.
	for i := 0; i < 10; i++ {
		println(i)
	}

	// Exemplo 2: Loop estilo "while".
	// Continua enquanto a condição (i < 10) for verdadeira.
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// Exemplo 3: Loop infinito.
	// O "break" garante que o loop não seja inatingível.
	for {
		println("Infinite loop")
		break // Para sair do loop
	}

	// Exemplo 4: Loop sobre um array/slice.
	// Definindo um slice de inteiros.
	numbers := []int{1, 2, 3, 4, 5}

	// Loop tradicional para percorrer o slice.
	for i := 0; i < len(numbers); i++ {
		println(numbers[i])
	}

	// Exemplo 5: Loop "for range" para percorrer slices.
	for index, value := range numbers {
		println("Index:", index, "Value:", value)
	}

	// Exemplo 6: Loop sobre mapas (map).
	// Definindo um mapa com chaves e valores.
	personAge := map[string]int{
		"Ana":   28,
		"João":  34,
		"Maria": 22,
	}

	// Usando o loop "for range" para percorrer o mapa.
	for name, age := range personAge {
		println("Name:", name, "Age:", age)
	}

	// Exemplo 7: Loop com break.
	// Saindo do loop quando uma condição for satisfeita.
	for i := 0; i < 10; i++ {
		if i == 5 {
			println("Breaking the loop at i =", i)
			break
		}
		println(i)
	}

	// Exemplo 8: Loop com continue.
	// Pulando para a próxima iteração quando a condição é verdadeira.
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Pula se i for par
		}
		println(i) // Só imprime valores ímpares
	}

	// Exemplo 9: Loop aninhado (Nested loop).
	// Loop externo e interno.
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			println("i:", i, "j:", j)
		}
	}

	// Exemplo 10: Simulação de loop "do-while".
	// Continua até que a condição seja satisfeita.
	i = 0
	for {
		println(i)
		i++
		if i >= 5 {
			break // Sai do loop após 5 iterações
		}
	}

	// Exemplo 11: Loop com múltiplas variáveis.
	// Incrementando duas variáveis ao mesmo tempo.
	for i, j := 0, 10; i < 5 && j > 5; i, j = i+1, j-1 {
		println("i:", i, "j:", j)
	}
}
