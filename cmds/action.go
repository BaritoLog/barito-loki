package cmds

import (
	"fmt"

	"github.com/BaritoLog/go-boilerplate/srvkit"
	"github.com/urfave/cli"
	"github.com/vwidjaya/barito-loki/loki"
)

func ActionBaritoLokiService(c *cli.Context) (err error) {
	serviceParams := map[string]interface{}{
		"grpcAddr":  configGrpcAddress(),
		"lokiUrl":      configLokiUrl(),
		"batchWaitMs":  configLokiBatchWaitMs(),
		"batchSize":    configLokiBatchSize(),
		"minBackoffMs": configLokiMinBackoffMs(),
		"maxBackoffMs": configLokiMaxBackoffMs(),
		"maxRetries":   configLokiMaxRetries(),
		"timeoutMs":    configLokiTimeoutMs(),
	}

	service, err := loki.NewBaritoLokiService(serviceParams)
	if err != nil {
		return
	}

	err = service.Start()
	if err != nil {
		return
	}
	fmt.Println("Loki's Promtail client started.")
	srvkit.AsyncGracefulShutdown(service.Close)

	return
}
