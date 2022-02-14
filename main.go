package main

import (
	"log"
	"os"

	"github.com/duongnln96/go-grpc-practice/cmd/client"
	"github.com/duongnln96/go-grpc-practice/cmd/server"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Commands = []*cli.Command{
		{
			Name:    "server",
			Usage:   "start server",
			Aliases: []string{"s"},
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:    "port",
					Aliases: []string{"p"},
					Value:   8080,
					Usage:   "start server with port",
				},
			},
			Action: func(c *cli.Context) error {
				return server.Start(c)
			},
		},
		{
			Name:    "client",
			Usage:   "start client",
			Aliases: []string{"c"},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "address",
					Aliases: []string{"addr"},
					Value:   "0.0.0.0",
					Usage:   "connect client to server address",
				},
				&cli.StringFlag{
					Name:    "port",
					Aliases: []string{"p"},
					Value:   "8080",
					Usage:   "connect to server port",
				},
			},
			Action: func(c *cli.Context) error {
				return client.Start(c)
			},
		},
	}

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "s")
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
