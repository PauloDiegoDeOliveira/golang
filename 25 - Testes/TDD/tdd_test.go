package tdd

import "testing"

// Testando a função Soma usando TDD
func TestSoma(t *testing.T) {
	// Casos de teste que queremos validar
	cases := []struct {
		name     string
		numeros  []int // Números que queremos somar
		esperado int   // O resultado esperado da soma
	}{
		{"Soma de dois números", []int{2, 3}, 5},    // 2 + 3 = 5
		{"Soma de três números", []int{1, 2, 3}, 6}, // 1 + 2 + 3 = 6
	}

	// Itera sobre os casos de teste
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Executa a função Soma (que ainda não existe, e deve falhar)
			resultado := Soma(c.numeros)

			// Verifica se o resultado é o esperado
			if resultado != c.esperado {
				t.Errorf("Erro no teste %s: esperado %d, mas obteve %d", c.name, c.esperado, resultado)
			}
		})
	}
}
