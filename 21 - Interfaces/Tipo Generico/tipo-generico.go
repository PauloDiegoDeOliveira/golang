package main

import "fmt"

// Definindo uma estrutura Stack que armazena elementos do tipo interface{}
// interface{} permite armazenar qualquer tipo, mas a manipulação dos dados
// exige conversão explícita para o tipo correto ao desempilhar
type Stack struct {
	elements []interface{}
}

// Método para empilhar (adicionar um elemento à Stack)
func (s *Stack) Push(value interface{}) {
	s.elements = append(s.elements, value)
}

// Método para desempilhar (remover o último elemento da Stack)
func (s *Stack) Pop() (interface{}, bool) {
	if len(s.elements) == 0 {
		return nil, false // Retorna nil e falso se a pilha estiver vazia
	}
	lastIndex := len(s.elements) - 1
	value := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return value, true
}

func main() {
	// Criando uma Stack
	var stack Stack

	// Empilhando valores inteiros
	stack.Push(10)
	stack.Push(20)

	// Desempilhando e convertendo para int
	value, ok := stack.Pop()
	if ok {
		// O valor retornado é do tipo interface{}, então precisa ser convertido para int
		intValue := value.(int)
		fmt.Println("Desempilhado:", intValue) // Saída: Desempilhado: 20
	}

	// Empilhando valores string
	stack.Push("Golang")
	stack.Push("Genéricos")

	// Interface baseada em tipo genérico
	// Isso pode gerar confusão, pois permite qualquer tipo de chave e valor.
	// Não há segurança de tipo em tempo de compilação, o que significa que erros só serão detectados em tempo de execução.
	// Além disso, exige conversões explícitas ao acessar os elementos.
	mapa := map[interface{}]interface{}{
		1:            "um",        // Inteiro como chave, string como valor
		float32(100): true,        // float32 como chave, bool como valor
		"3":          "três",      // String como chave e valor
		true:         float32(12), // Bool como chave, float32 como valor
	}

	// Exemplo confuso, pois misturamos vários tipos no mapa sem nenhuma estrutura clara
	fmt.Println("Mapa confuso:", mapa)

	// Para acessar os valores, precisamos converter explicitamente os tipos
	if val, ok := mapa[1]; ok {
		// Precisamos verificar o tipo e fazer a conversão
		if strVal, ok := val.(string); ok {
			fmt.Println("Valor para a chave 1:", strVal)
		} else {
			fmt.Println("O valor para a chave 1 não é uma string")
		}
	}

	// Acessando a chave float32(100) e tentando converter o valor para bool
	if val, ok := mapa[float32(100)]; ok {
		// Novamente, precisamos verificar o tipo e converter explicitamente
		if boolVal, ok := val.(bool); ok {
			fmt.Println("Valor para a chave float32(100):", boolVal)
		} else {
			fmt.Println("O valor para a chave float32(100) não é um bool")
		}
	}

	// Tentando acessar a chave "3" e convertendo para string
	if val, ok := mapa["3"]; ok {
		if strVal, ok := val.(string); ok {
			fmt.Println("Valor para a chave '3':", strVal)
		} else {
			fmt.Println("O valor para a chave '3' não é uma string")
		}
	}

	// Tentando acessar a chave true e convertendo para float32
	if val, ok := mapa[true]; ok {
		if floatVal, ok := val.(float32); ok {
			fmt.Println("Valor para a chave true:", floatVal)
		} else {
			fmt.Println("O valor para a chave true não é um float32")
		}
	}
}
