package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Ip Finder"
	app.Usage = "Aplicação pra busca de ip e nomes de servidores na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com",
		},
	}

	app.Commands =[]cli.Command{
		{
			Name: "ip",
			Usage: "Busca ips de endereços na internet",
			Flags: flags,
			Action: buscarIps,
		},

		{
			Name: "servidores",
			Usage: "Busca o nome do servidor na internet",
			Flags: flags,
			Action: buscarServidor,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func buscarServidor(c *cli.Context) {
	host := c.String("host")
	
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}