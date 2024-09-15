package formas

import (
	"math"
	"testing"
)

// Testando a área do Retângulo e Círculo com subtestes usando t.Run
func TestArea(t *testing.T) {
	// Testando o cálculo da área para Retângulo
	t.Run("Área do Retângulo", func(t *testing.T) {
		ret := Retangulo{largura: 10, altura: 5}
		esperado := 50.0
		resultado := ret.area()

		if resultado != esperado {
			t.Errorf("Esperado %f, mas obteve %f", esperado, resultado)
		}
	})

	// Testando o cálculo da área para Círculo
	t.Run("Área do Círculo", func(t *testing.T) {
		circ := Circulo{raio: 7}
		esperado := math.Pi * circ.raio * circ.raio
		resultado := circ.area()

		if resultado != esperado {
			t.Errorf("Esperado %f, mas obteve %f", esperado, resultado)
		}
	})
}

// Testando a função areaTotal para Retângulo e Círculo com subtestes usando t.Run
func TestAreaTotal(t *testing.T) {
	// Testando areaTotal com um Retângulo
	t.Run("Área Total do Retângulo", func(t *testing.T) {
		ret := Retangulo{largura: 10, altura: 5}
		esperado := 50.0
		resultado := areaTotal(ret)

		if resultado != esperado {
			t.Errorf("Esperado %f, mas obteve %f", esperado, resultado)
		}
	})

	// Testando areaTotal com um Círculo
	t.Run("Área Total do Círculo", func(t *testing.T) {
		circ := Circulo{raio: 7}
		esperado := math.Pi * circ.raio * circ.raio
		resultado := areaTotal(circ)

		if resultado != esperado {
			t.Errorf("Esperado %f, mas obteve %f", esperado, resultado)
		}
	})
}
