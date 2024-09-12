package app

// Importamos os pacotes necessários.
import (
	"fmt"
	"net"

	"github.com/urfave/cli/v2" // Importa a versão 2 do pacote urfave/cli
)

// A função Gerar cria e retorna uma aplicação CLI configurada. Essa função será
// chamada no arquivo principal para iniciar a aplicação.
func Gerar() *cli.App {
	// Definimos as flags fora da estrutura do comando.
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "host",                    // Nome da flag (usado como --host)
			Aliases: []string{"H"},             // Alias da flag.
			Value:   "google.com",              // Valor padrão da flag.
			Usage:   "Domínio para buscar IPs", // Descrição da flag.
		},
	}

	// Criamos uma nova instância de `cli.App`, agora usando a versão 2.
	app := &cli.App{
		// Definimos algumas propriedades básicas da aplicação CLI.
		Name:    "Aplicação de Linha de Comando",            // Nome da aplicação
		Usage:   "Busca IPs e Nomes de Domínio na Internet", // Propósito da aplicação
		Version: "1.0.0",                                    // Versão da aplicação

		// Aqui definimos os comandos que nossa aplicação pode executar.
		Commands: []*cli.Command{
			{
				// Nome do comando que o usuário pode executar.
				Name:  "ip",
				Usage: "Busca IPs de endereços na internet",
				Flags: flags, // Atribuímos as flags definidas fora do comando.
				// A função `Action` define o que será executado quando o comando for chamado.
				Action: func(c *cli.Context) error {
					buscarIps(c) // Chama a função buscarIps, que busca os IPs de um domínio.
					return nil
				},
			},
			{
				// Nome do comando que o usuário pode executar.
				Name:  "servidores",
				Usage: "Busca o nome do Servidores na Internet",
				Flags: flags, // Atribuímos as flags definidas fora do comando.
				// A função `Action` define o que será executado quando o comando for chamado.
				Action: func(c *cli.Context) error {
					buscarServidores(c) // Chama a função buscarIps, que busca os IPs de um domínio.
					return nil
				},
			},
		},
	}

	// A função retorna a aplicação configurada.
	return app
}

// Função responsável por buscar os endereços IP do domínio fornecido.
func buscarIps(c *cli.Context) {
	// Obtém o valor da flag "host" passada pelo usuário ou o valor padrão.
	host := c.String("host")

	// Realiza a busca dos endereços IP do domínio fornecido.
	ips, err := net.LookupIP(host)

	// Se ocorrer um erro ao buscar os IPs, exibe uma mensagem de erro no terminal.
	if err != nil {
		fmt.Println("Erro ao buscar IPs:", err)
		return // Encerra a função se houver erro.
	}

	// Se a busca for bem-sucedida, percorre todos os IPs encontrados e os imprime no terminal.
	for _, ip := range ips {
		fmt.Println(ip) // Exibe o IP no terminal.
	}
}

// Função responsável por buscar os endereços IP do domínio fornecido.
func buscarServidores(c *cli.Context) {
	// Obtém o valor da flag "host" passada pelo usuário ou o valor padrão.
	host := c.String("host")

	// Realiza a busca dos servidores fornecido.
	servidores, err := net.LookupNS(host)

	// Se ocorrer um erro ao buscar os servidores, exibe uma mensagem de erro no terminal.
	if err != nil {
		fmt.Println("Erro ao buscar Servidores:", err)
		return // Encerra a função se houver erro.
	}

	// Se a busca for bem-sucedida, percorre todos os servidores encontrados e os imprime no terminal.
	for _, ip := range servidores {
		fmt.Println(ip.Host) // Exibe o Servidor no terminal.
	}
}
