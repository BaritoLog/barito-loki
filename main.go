package main

import (
	"fmt"
	"os"

	"github.com/BaritoLog/barito-loki/loki"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const (
	Name    = "barito-loki"
	Version = "0.1"
)

var (
	Commit string = "N/A"
	Build  string = "MANUAL"
)

func init() {
	log.SetLevel(log.WarnLevel)
}

func main() {
	app := cli.App{
		Name:    Name,
		Usage:   "Forward logs to Loki for Barito project",
		Version: fmt.Sprintf("%s-%s-%s", Version, Build, Commit),
		Action:  loki.Start,
		Before: func(c *cli.Context) error {
			fmt.Fprintf(os.Stderr, "%s Started. Version: %s Build: %s Commit: %s\n", Name, Version, Build, Commit)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Some error occurred: %s", err.Error())
	}
}
