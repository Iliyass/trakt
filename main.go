package main

import (
	"fmt"
	"log"
	"os"

	trakt "github.com/Iliyass/trakt/Trakt"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Trakt"
	app.Usage = "Manage your time"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Ok Cool!")
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "tags",
			Aliases: []string{"t"},
			Usage:   "tags management",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add new Tag",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "data, d", Required: true},
					},
					Action: func(c *cli.Context) error {
						data := []byte(c.String("data"))
						_, err := trakt.AddTag(data)
						if err != nil {
							panic(err)
						}
						fmt.Println("Success! Tag has been added")
						return nil
					},
				},
			},
		},
		{
			Name:    "trakt",
			Aliases: []string{"t"},
			Usage:   "trakt management",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add new Trakt",
					Flags: []cli.Flag{
						cli.StringFlag{Name: "data, d", Required: true},
					},
					Action: func(c *cli.Context) error {
						data := []byte(c.String("data"))
						_, err := trakt.AddTrakt(data)
						if err != nil {
							panic(err)
						}
						fmt.Println("Success! Trakt has been added")
						return nil
					},
				},
				{
					Name:  "list",
					Usage: "list Trakt",
					Flags: []cli.Flag{
						cli.Int64Flag{Name: "from, f", Required: true},
						cli.Int64Flag{Name: "to, t", Required: true},
					},
					Action: func(c *cli.Context) error {
						from := c.Int64("from")
						to := c.Int64("to")
						trakts := trakt.GetTraktsByDate(from, to)
						fmt.Println(string(trakts))
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
