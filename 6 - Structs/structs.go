package main

import (
	"fmt"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Age       int
	Endereco  Endereco
}

type Endereco struct {
	Cidade string
	Estado string
	Pais   string
	Cep    string
	Numero int
	Compl  string
}

func main() {
	fmt.Println("Structs")

	var u User
	u.FirstName = "Paulo"
	u.LastName = "Diego"
	u.Email = "teste@teste.com"

	fmt.Println(u)

	u2 := User{
		FirstName: "Paulo",
		LastName:  "Diego",
		Email:     "teste@teste.com",
		Age:       30,
		Endereco: Endereco{
			Cidade: "São Paulo",
			Estado: "SP",
			Pais:   "Brasil",
			Cep:    "00000-000",
			Numero: 123,
			Compl:  "Apto 123",
		},
	}

	fmt.Println(u2, u2.Endereco.Cidade)

	usuario3 := User{FirstName: "Paulo"}
	fmt.Println(usuario3)
}
