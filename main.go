package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "save2cloud",
		Description:          "Backup a folder to any cloud using the CLI",
		Version:              "v0.1",
		Authors:              []*cli.Author{{Name: "Sean Hellum"}},
		EnableBashCompletion: true,
		Action: func(c *cli.Context) error {
			provider := c.String("provider")
			if provider == "gcp" {
				GcpEntrypoint(c)
			}
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "provider",
				Value: "gcp",
				Usage: "Cloud provider to save the file to. Options: gcp",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		PrintError("Unknown error has occurred.", 1)
	}
}
