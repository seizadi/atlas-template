package commands

import (
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	{
		Name:   "create",
		Usage:  "create",
		Action: createTemplate,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "path, p",
				Usage: "path to the template files",
			},
			cli.StringFlag{
				Name:  "repo, r",
				Usage: "path to the resolved files",
			},
		},
	},
}