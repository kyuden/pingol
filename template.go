package main

const AppHelpTemplate string = `NAME:
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
