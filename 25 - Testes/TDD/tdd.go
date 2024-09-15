package tdd

// Função Soma que recebe um slice de números e retorna a soma deles
func Soma(numeros []int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
