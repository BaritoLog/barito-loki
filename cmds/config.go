package cmds

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

const (
	EnvServiceAddress = "BARITO_LOKI_SERVICE_ADDRESS"
	EnvConsulUrl      = "BARITO_CONSUL_URL"
	EnvConsulLokiName = "BARITO_CONSUL_LOKI_NAME"

	EnvLokiUrl         = "BARITO_LOKI_URL"
	EnvLokiBatchSize   = "BARITO_LOKI_BATCH_SIZE"
	EnvLokiBatchWaitMs = "BARITO_LOKI_BATCH_WAIT_MS"

	EnvLokiMinBackoffMs = "BARITO_LOKI_MIN_BACKOFF_MS"
	EnvLokiMaxBackoffMs = "BARITO_LOKI_MAX_BACKOFF_MS"
	EnvLokiMaxRetries   = "BARITO_LOKI_MAX_RETRIES"
	EnvLokiTimeoutMs    = "BARITO_LOKI_TIMEOUT_MS"
)

var (
	DefaultServiceAddress = ":8080"
	DefaultConsulLokiName = "loki"

	DefaultLokiUrl         = "http://localhost:3100"
	DefaultLokiBatchSize   = 50000
	DefaultLokiBatchWaitMs = 500

	DefaultLokiMinBackoffMs = 100
	DefaultLokiMaxBackoffMs = 10000
	DefaultLokiMaxRetries   = 10
	DefaultLokiTimeoutMs    = 10000
)

func configServiceAddress() (s string) {
	return stringEnvOrDefault(EnvServiceAddress, DefaultServiceAddress)
}

func configConsulUrl() (s string) {
	return os.Getenv(EnvConsulUrl)
}

func configConsulLokiName() (s string) {
	return stringEnvOrDefault(EnvConsulLokiName, DefaultConsulLokiName)
}

func configLokiUrl() (url string) {
	consulUrl := configConsulUrl()
	name := configConsulLokiName()
	url, err := consulLokiUrl(consulUrl, name)
	if err != nil {
		url = stringEnvOrDefault(EnvLokiUrl, DefaultLokiUrl)
		return
	}

	logConfig("consul", EnvLokiUrl, url)
	return
}

func configLokiBatchSize() (i int) {
	return intEnvOrDefault(EnvLokiBatchSize, DefaultLokiBatchSize)
}

func configLokiBatchWaitMs() (i int) {
	return intEnvOrDefault(EnvLokiBatchWaitMs, DefaultLokiBatchWaitMs)
}

func configLokiMinBackoffMs() (i int) {
	return intEnvOrDefault(EnvLokiMinBackoffMs, DefaultLokiMinBackoffMs)
}

func configLokiMaxBackoffMs() (i int) {
	return intEnvOrDefault(EnvLokiMaxBackoffMs, DefaultLokiMaxBackoffMs)
}

func configLokiMaxRetries() (i int) {
	return intEnvOrDefault(EnvLokiMaxRetries, DefaultLokiMaxRetries)
}

func configLokiTimeoutMs() (i int) {
	return intEnvOrDefault(EnvLokiTimeoutMs, DefaultLokiTimeoutMs)
}

func stringEnvOrDefault(key, defaultValue string) string {
	s := os.Getenv(key)
	if len(s) > 0 {
		logConfig("env", key, s)
		return s
	}

	logConfig("default", key, defaultValue)
	return defaultValue
}

func intEnvOrDefault(key string, defaultValue int) int {
	s := os.Getenv(key)
	i, err := strconv.Atoi(s)
	if err == nil {
		logConfig("env", key, i)
		return i
	}

	logConfig("default", key, defaultValue)
	return defaultValue
}

func logConfig(source, key string, val interface{}) {
	log.WithField("config", source).Warnf("%s = %v", key, val)
}
