package main

import "github.com/codegangsta/cli"

var (
	commands = []cli.Command{
		{
			Name:      "add",
			ShortName: "a",
			Usage:     "Add new or update environment variable",
			Flags: []cli.Flag{
				flKey,
				flValue,
			},
			Action: addCommand,
		},
		{
			Name:      "print",
			ShortName: "p",
			Usage:     "Prints the stored environment variables",
			Action:    printCommand,
		},
		{
			Name:            "run",
			ShortName:       "r",
			Usage:           "Runs the specified command with stored environments",
			SkipFlagParsing: true,
			Action:          runCommand,
		},
	}
)
