package main

import (
	"fmt"
)

func main() {
	// Exemplo 1: Map simples com chave e valor de tipos primitivos
	m := make(map[string]string)

	// Adicionando pares chave/valor ao map
	m["name"] = "Paulo"
	m["age"] = "30"
	m["city"] = "New York"

	// Acessando um valor no map
	fmt.Println("Nome:", m["name"])

	// Verificando se uma chave existe no map
	if age, exists := m["age"]; exists {
		fmt.Println("Idade:", age)
	} else {
		fmt.Println("A chave 'age' não existe.")
	}

	// Deletando uma chave
	delete(m, "city")

	// Iterando sobre o map
	for key, value := range m {
		fmt.Printf("Chave: %s, Valor: %s\n", key, value)
	}

	// Exemplo 2: Map inicializado com valores
	usuarios := map[string]string{
		"nome":   "Alice",
		"idade":  "25",
		"cidade": "São Paulo",
	}

	fmt.Println(usuarios)
	fmt.Println("Nome:", usuarios["nome"])

	// Exemplo 3: Map aninhado
	usuario := map[string]map[string]string{
		"perfil": {
			"nome":   "Bob",
			"idade":  "22",
			"cidade": "Rio de Janeiro",
		},
	}

	// Acessando o map aninhado
	fmt.Println("Nome:", usuario["perfil"]["nome"])
	fmt.Println("Cidade:", usuario["perfil"]["cidade"])

	// Exemplo 4: Verificação da existência de uma chave
	usuarios2 := map[string]string{
		"nome":   "Carlos",
		"idade":  "40",
		"cidade": "Brasília",
	}

	// Verificando se a chave "email" existe
	if email, ok := usuarios2["email"]; ok {
		fmt.Println("Email:", email)
	} else {
		fmt.Println("Chave 'email' não existe no map.")
	}

	// Exemplo 5: Map com tipos numéricos
	precos := map[int]float64{
		1: 10.50,
		2: 15.75,
		3: 20.30,
	}

	// Iterando sobre o map
	for produto, preco := range precos {
		fmt.Printf("Produto %d: R$ %.2f\n", produto, preco)
	}

	// Exemplo 6: Map com slice como valor
	usuario3 := map[string][]string{
		"nome": {"Alice", "Bob", "Carlos"},
	}

	// Acessando e imprimindo o slice dentro do map
	for _, nome := range usuario3["nome"] {
		fmt.Println("Nome:", nome)
	}

	// Exemplo 7: Manipulando maps com len (tamanho do map)
	usuarios3 := map[string]string{
		"nome":   "David",
		"idade":  "28",
		"cidade": "Curitiba",
	}

	// Imprimindo o tamanho do map
	fmt.Println("Tamanho do map:", len(usuarios3))
}
