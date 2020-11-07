package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate Return the app
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line app"
	app.Usage = "Search Ip's and names of servers on the web"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IP's on the web",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servers",
			Usage:  "Serch servers names on web",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
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

	for _, ns := range servers {
		fmt.Println(ns)
	}
}
