package enderecos

import (
	"testing" // Pacote de testes padrão do Go
)

// Função de teste para TipoEndereco
func TestTipoEndereco(t *testing.T) {
	// Casos de teste que queremos validar
	cases := []struct {
		input    string
		expected string
	}{
		{"Rua dos Bobos 0", "Rua"},           // Teste para "Rua"
		{"Avenida Paulista 123", "Avenida"},  // Teste para "Avenida"
		{"Alameda Santos 456", "Alameda"},    // Teste para "Alameda"
		{"Praça da Sé 789", "Tipo Inválido"}, // Teste para tipo inválido
	}

	// Itera sobre os casos de teste
	for _, c := range cases {
		// Executa a função TipoEndereco com o endereço de entrada
		result := TipoEndereco(c.input)

		// Verifica se o resultado da função é o esperado
		if result != c.expected {
			// Caso contrário, falha o teste
			t.Errorf("TipoEndereco(%q) == %q, esperado %q", c.input, result, c.expected)
		}
	}
}
