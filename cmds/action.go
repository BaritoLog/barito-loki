package cmds

import (
	"fmt"

	"github.com/BaritoLog/go-boilerplate/srvkit"
	"github.com/urfave/cli"
	"github.com/vwidjaya/barito-loki/loki"
)

func ActionBaritoLokiService(c *cli.Context) (err error) {
	serviceParams := map[string]interface{}{
		"grpcAddr":     configGrpcAddress(),
		"restAddr":     configRestAddress(),
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

	go service.Start()
	if configServeRestApi() {
		go service.LaunchREST()
	}

	fmt.Println("Barito-Loki started.")
	srvkit.GracefullShutdown(service.Close)

	return
}
