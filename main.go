package main

import (
	"github.com/urfave/cli/v2"
	"github.com/shyim/ssh-host-manager/cmd"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name: "SSH Host Manager",
		Version: "0.1.0",
		EnableBashCompletion: true,
		Usage: "test",
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List all entries",
				Action: cmd.ListCommand,
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Add a new entry",
				Action: cmd.AddCommand,
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "Delete a entry",
				Action: cmd.DeleteCommand,
			},
			{
				Name:    "edit",
				Aliases: []string{"e"},
				Usage:   "Edit a entry",
				Action: cmd.EditCommand,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
