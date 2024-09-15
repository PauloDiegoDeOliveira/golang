package enderecos

import (
	"testing" // Pacote de testes padrão do Go
)

// Função de teste para TipoEndereco usando subtestes para cada cenário
func TestTipoEndereco(t *testing.T) {
	t.Parallel() // Indica que os subtestes podem ser executados em paralelo

	// Casos de teste que queremos validar
	cases := []struct {
		name     string // Nome do cenário de teste
		input    string // Endereço de entrada
		expected string // Resultado esperado
	}{
		{"Rua", "Rua dos Bobos 0", "Rua"},                              // Teste para "Rua"
		{"Avenida", "Avenida Paulista 123", "Avenida"},                 // Teste para "Avenida"
		{"Alameda", "Alameda Santos 456", "Alameda"},                   // Teste para "Alameda"
		{"Tipo Inválido", "Praça da Sé 789", "Tipo Inválido"},          // Teste para tipo inválido
		{"Outro Tipo Inválido", "Estrada do Sol 999", "Tipo Inválido"}, // Outro teste para tipo inválido
	}

	// Itera sobre os casos de teste
	for _, c := range cases {
		// Executa cada cenário como um subteste
		t.Run(c.name, func(t *testing.T) {
			// Executa a função TipoEndereco com o endereço de entrada
			result := TipoEndereco(c.input)

			// Verifica se o resultado da função é o esperado
			if result != c.expected {
				// Caso contrário, falha o teste
				t.Errorf("TipoEndereco(%q) == %q, esperado %q", c.input, result, c.expected)
			}
		})
	}
}
