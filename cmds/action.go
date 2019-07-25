package cmds

import (
	"fmt"

	"github.com/BaritoLog/go-boilerplate/srvkit"
	"github.com/urfave/cli"
	"github.com/vwidjaya/barito-loki/loki"
)

func ActionBaritoLokiService(c *cli.Context) (err error) {
	address := configServiceAddress()
	lokiUrl := configLokiUrl()
	bulkSize := configLokiBulkSize()
	flushMs := configLokiFlushIntervalMs()

	lkConfig := loki.NewLokiConfig(lokiUrl, bulkSize, flushMs)
	service := loki.NewBaritoLokiService(address, lkConfig)

	err = service.Start()
	if err != nil {
		return
	}
	fmt.Println("Loki client started.")
	srvkit.AsyncGracefulShutdown(service.Close)

	return
}
