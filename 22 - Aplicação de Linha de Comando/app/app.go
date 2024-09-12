package app

import "github.com/urfave/cli"

// Gerar - Função para gerar a aplicação de linha de comando
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Domínio na Internet"
	app.Version = "1.0.0"

	return app
}
