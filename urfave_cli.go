package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "name",
			Aliases: []string{"c"},
			Usage:   "Load configuration from `FILE`",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:  "list",
			Usage: "List students",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "json",
					Usage: "output as JSON",
					Value: false,
				},
			},
			Action: cmdList,
		},
	}
	app.Name = "score"
	app.Usage = "Show student's score"
	app.Run(os.Args)
}

func cmdList(c *cli.Context) error {
	if c.Bool("json") {
		return nil
	}
	return nil
}
