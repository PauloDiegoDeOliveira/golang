package main

import (
	"fmt"
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
// Agora 'Retangulo' implementa a interface 'Forma'
func (r Retangulo) area() float64 {
	return r.largura * r.altura
}

// Definindo o struct 'Circulo'
type Circulo struct {
	raio float64
}

// Implementação do método 'area' para o struct 'Circulo'
// Agora 'Circulo' também implementa a interface 'Forma'
func (c Circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

// Função que calcula a área total de qualquer objeto que implemente a interface 'Forma'
func areaTotal(f Forma) float64 {
	return f.area()
}

func main() {
	// Criando uma instância de 'Retangulo'
	ret := Retangulo{largura: 10, altura: 5}

	// Criando uma instância de 'Circulo'
	circ := Circulo{raio: 7}

	// Calculando a área de cada forma individualmente usando a função 'areaTotal'
	fmt.Println("Área do retângulo:", areaTotal(ret)) // Saída: 50
	fmt.Println("Área do círculo:", areaTotal(circ))  // Saída: 153.93804002589985
}
