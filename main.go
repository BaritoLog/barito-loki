package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vwidjaya/barito-loki/cmds"
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
		Action:  cmds.ActionBaritoLokiService,
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
