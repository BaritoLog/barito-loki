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

func TestGetConsulLokiName(t *testing.T) {
	FatalIf(t, configConsulLokiName() != DefaultConsulLokiName, "should return default ")

	os.Setenv(EnvConsulLokiName, "thor")
	defer os.Clearenv()
	FatalIf(t, configConsulLokiName() != "thor", "should get from env variable")
}

func TestGetLokiUrl(t *testing.T) {
	FatalIf(t, configLokiUrl() != DefaultLokiUrl, "should return default ")

	os.Setenv(EnvLokiUrl, "http://some-loki")
	defer os.Clearenv()
	FatalIf(t, configLokiUrl() != "http://some-loki", "should get from env variable")
}

func TestGetLokiBatchSize(t *testing.T) {
	FatalIf(t, configLokiBatchSize() != DefaultLokiBatchSize, "should return default ")

	os.Setenv(EnvLokiBatchSize, "999")
	defer os.Clearenv()
	FatalIf(t, configLokiBatchSize() != 999, "should get from env variable")
}

func TestGetLokiBatchWaitMs(t *testing.T) {
	FatalIf(t, configLokiBatchWaitMs() != DefaultLokiBatchWaitMs, "should return default ")

	os.Setenv(EnvLokiBatchWaitMs, "999")
	defer os.Clearenv()
	FatalIf(t, configLokiBatchWaitMs() != 999, "should get from env variable")
}

func TestGetLokiMinBackoffMs(t *testing.T) {
	FatalIf(t, configLokiMinBackoffMs() != DefaultLokiMinBackoffMs, "should return default ")

	os.Setenv(EnvLokiMinBackoffMs, "999")
	defer os.Clearenv()
	FatalIf(t, configLokiMinBackoffMs() != 999, "should get from env variable")
}

func TestGetLokiMaxBackoffMs(t *testing.T) {
	FatalIf(t, configLokiMaxBackoffMs() != DefaultLokiMaxBackoffMs, "should return default ")

	os.Setenv(EnvLokiMaxBackoffMs, "999")
	defer os.Clearenv()
	FatalIf(t, configLokiMaxBackoffMs() != 999, "should get from env variable")
}

func TestGetLokiMaxRetries(t *testing.T) {
	FatalIf(t, configLokiMaxRetries() != DefaultLokiMaxRetries, "should return default ")

	os.Setenv(EnvLokiMaxRetries, "999")
	defer os.Clearenv()
	FatalIf(t, configLokiMaxRetries() != 999, "should get from env variable")
}

func TestGetLokiTimeoutMs(t *testing.T) {
	FatalIf(t, configLokiTimeoutMs() != DefaultLokiTimeoutMs, "should return default ")

	os.Setenv(EnvLokiTimeoutMs, "999")
	defer os.Clearenv()
	FatalIf(t, configLokiTimeoutMs() != 999, "should get from env variable")
}
