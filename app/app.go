package app

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Ip Finder CLI"
	app.Usage = "App to find ips and servers name"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "",
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


	if host == "" {
		ip := getPublicAddress()
		fmt.Println(ip)

		return
	}
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func searchServers(c *cli.Context) {
	host := c.String("host")
	
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

func getFirstNonLoopbackIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}
	return ""
}

func getPublicAddress() string {
	resp, err := http.Get("https://ifconfig.me")

	if err != nil {
		fmt.Printf("Error: %v", err.Error())
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error: %v", err.Error())
	}

	return string(data)

}
