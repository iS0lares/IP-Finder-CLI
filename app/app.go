package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Ip Finder CLI"
	app.Usage = "App to find ips and servers name"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com",
		},
	}

	app.Commands =[]cli.Command{
		{
			Name: "ip",
			Usage: "search ip address on internet",
			Flags: flags,
			Action: searchIp,
		},

		{
			Name: "server",
			Usage: "Search name of servers on internet",
			Flags: flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIp(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func searchServers(c *cli.Context) {
	host := c.String("host")
	
	servers, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}