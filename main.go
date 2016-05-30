package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.HideHelp = true
	app.Author = "kyuden"
	app.Usage = "A simple CLI tool for finding your iPhone"

	if len(os.Args) > 1 {
		app.Action = func(c *cli.Context) error {
			fmt.Println("Hello", c.Args()[0])
			return nil
		}
	}

	cli.AppHelpTemplate = AppHelpTemplate

	app.Run(os.Args)
}
