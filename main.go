package main

import (
	"log"
	"os"
	"sort"
	"github.com/exuan/waka-api/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		UseShortOptionHandling: true,
		Name:                   "waka-api service",
		Usage:                  "",
		Commands: cli.Commands{
			cmd.Server,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
