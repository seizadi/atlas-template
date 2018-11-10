package main

import (
	"fmt"
	"os"
	"log"
	
	"github.com/seizadi/atlas-template/cmd/atlas-template/commands"
	"github.com/codegangsta/cli"
)

var version = "0.1"

func main() {
	app := cli.NewApp()
	app.Name = "Template Engine CLI"
	app.Usage = ""
	app.Author = "Soheil Eizadi"
	app.Email = "seizadi@gmail.com"
	app.Version = version
	app.Flags = []cli.Flag{}
	
	app.Before = func(c *cli.Context) error {
		return nil
	}
	app.Commands = commands.Commands
	app.CommandNotFound = cmdNotFound
	
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("%v", err)
	}
}


func cmdNotFound(c *cli.Context, command string) {
	fmt.Printf(
		"%s: '%s' is not a %s command. See '%s --help'.\n",
		c.App.Name,
		command,
		c.App.Name,
		os.Args[0],
	)
	os.Exit(1)
}