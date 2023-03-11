package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func Run() error {
	app := &cli.App{
		Name:  "myapp",
		Usage: "my awesome app",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config-file",
				Value:   "./config/example.yaml",
				Usage:   "path to configuration file",
				EnvVars: []string{"CONFIG_FILE"},
			},
		},
		Action: func(c *cli.Context) error {
			configFile := c.String("config-file")
			if configFile == "" {
				return cli.Exit("missing configuration file path", 1)
			}

			// TODO: call your application entry point function here
			log.Printf("config file path: %s\n", configFile)

			return nil
		},
	}

	return app.Run(os.Args)
}
