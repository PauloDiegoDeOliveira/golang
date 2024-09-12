package main

import (
	"linha-de-comando/app" // Pacote onde está a função 'Gerar', que gera nossa aplicação CLI.
	"log"                  // Pacote usado para logar erros ou outras mensagens.
	"os"                   // Pacote para acessar argumentos e funcionalidades do sistema operacional.
)

func main() {
	// Exibe uma mensagem de texto "Ponto de Partida" no terminal quando o programa começa.
	println("Ponto de Partida")

	// Aqui, chamamos a função Gerar do pacote "app", que retorna uma instância da nossa aplicação de linha de comando.
	aplicacao := app.Gerar()

	// Tentamos rodar a aplicação passando os argumentos da linha de comando.
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

	// Exibe uma mensagem de texto "Ponto de Chegada" no terminal ao final da execução do programa.
	println("Ponto de Chegada")
}
