package main

import (
	"github.com/fjah/seachad/core/web"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	app := &cli.App{
		Name:  "seachad",
		Usage: "A free, open-source and customisable donation site",
		Commands: []*cli.Command{
			{
				Name:    "serve",
				Aliases: []string{"s"},
				Usage:   "Starts a Seachad server",
				Action: func(c *cli.Context) error {
					addr := c.Args().Get(0)
					if addr == "" {
						addr = ":8080"
					} else if _, err := strconv.Atoi(addr); err == nil {
						addr = ":" + addr
					}

					server := web.Server{
						Address: addr,
					}
					server.Init()
					return server.Run()
				},
			},
			{
				Name:    "docs",
				Aliases: []string{"d"},
				Usage:   "Generates documentation",
				Action: func(c *cli.Context) error {
					return exec.Command("swag", "init", "--dir", "core/web").Run()
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
