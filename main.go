package main

import (
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

	cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} [iPhone model name]

EXAMPLE:
   # when iphone 5s
   {{.Name}} 5s

   # when iphone 6
   {{.Name}} 6

   # when iphone 6+
   {{.Name}} 6.Plus

VERSION:
   {{.Version}}{{if or .Author .Email}}

AUTHOR:{{if .Author}}
  {{.Author}}{{if .Email}} - <{{.Email}}>{{end}}{{else}}
  {{.Email}}{{end}}{{end}}

OPTIONS:
   --help, -h	show help
   {{range .Flags}}{{.}}
   {{end}}
`

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
