package formas

import (
	"math"
)

// Definindo a interface Forma, que contém o método 'area'
type Forma interface {
	area() float64
}

// Definindo o struct 'Retangulo'
type Retangulo struct {
	largura, altura float64
}

// Implementação do método 'area' para o struct 'Retangulo'
func (r Retangulo) area() float64 {
	return r.largura * r.altura
}

// Definindo o struct 'Circulo'
type Circulo struct {
	raio float64
}

// Implementação do método 'area' para o struct 'Circulo'
func (c Circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

// Função que calcula a área total de qualquer objeto que implemente a interface 'Forma'
func areaTotal(f Forma) float64 {
	return f.area()
}
