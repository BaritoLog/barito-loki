package cmds

import (
	"os"
	"testing"

	. "github.com/BaritoLog/go-boilerplate/testkit"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.ErrorLevel)
}

func TestGetServiceAddress(t *testing.T) {
	FatalIf(t, configServiceAddress() != DefaultServiceAddress, "should return default ")

	os.Setenv(EnvServiceAddress, ":12345")
	defer os.Clearenv()
	FatalIf(t, configServiceAddress() != ":12345", "should get from env variable")
}

func TestGetLokiUrl(t *testing.T) {
	FatalIf(t, configLokiUrl() != DefaultLokiUrl, "should return default ")

	os.Setenv(EnvLokiUrl, "http://some-loki")
	defer os.Clearenv()
	FatalIf(t, configLokiUrl() != "http://some-loki", "should get from env variable")
}

func TestGetLokiBulkSize(t *testing.T) {
	FatalIf(t, configLokiBulkSize() != DefaultLokiBulkSize, "should return default ")

	os.Setenv(EnvLokiBulkSize, "999")
	defer os.Clearenv()
	FatalIf(t, configLokiBulkSize() != 999, "should get from env variable")
}

func TestGetLokiFlushIntervalMs(t *testing.T) {
	FatalIf(t, configLokiFlushIntervalMs() != DefaultLokiFlushIntervalMs, "should return default ")

	os.Setenv(EnvLokiFlushIntervalMs, "999")
	defer os.Clearenv()
	FatalIf(t, configLokiFlushIntervalMs() != 999, "should get from env variable")
}
