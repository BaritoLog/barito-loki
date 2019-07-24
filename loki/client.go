package loki

import (
	"fmt"

	"github.com/urfave/cli"
)

func Start(c *cli.Context) (err error) {
	fmt.Println("Loki client started.")
	return
}
