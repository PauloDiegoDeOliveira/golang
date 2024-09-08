package main

import "fmt"

// Definimos um tipo chamado "pessoa"
type pessoa struct {
	Nome  string
	Idade int
}

// Definimos um tipo chamado "estudante" que "compoe" o tipo pessoa
type estudante struct {
	pessoa // Isso significa que estudante tem todos os campos de pessoa
	Curso  string
}

func main() {
	// Exemplo de composição em Go
	fmt.Println("Composição em Go")

	// Criamos uma instância de "pessoa"
	p1 := pessoa{"Paulo", 30}
	fmt.Println(p1)

	// Criamos uma instância de "estudante", que "inclui" a pessoa p1
	e1 := estudante{p1, "Programação"}
	// Podemos acessar os campos de pessoa diretamente de e1, por exemplo, e1.Nome ou e1.Idade
	fmt.Println(e1.Idade) // Saída: 30
	fmt.Println(e1.Nome)  // Saída: Paulo

	// Também podemos acessar o campo do tipo estudante
	fmt.Println(e1.Curso)
}
