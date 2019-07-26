package cmds

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

const (
	EnvServiceAddress      = "BARITO_LOKI_SERVICE_ADDRESS"
	EnvConsulUrl           = "BARITO_CONSUL_URL"
	EnvConsulLokiName      = "BARITO_CONSUL_LOKI_NAME"
	EnvLokiUrl             = "BARITO_LOKI_URL"
	EnvLokiBulkSize        = "BARITO_LOKI_BULK_SIZE"
	EnvLokiFlushIntervalMs = "BARITO_LOKI_FLUSH_INTERVAL_MS"
)

var (
	DefaultServiceAddress      = ":8080"
	DefaultConsulLokiName      = "loki"
	DefaultLokiUrl             = "http://localhost:3100"
	DefaultLokiBulkSize        = 500
	DefaultLokiFlushIntervalMs = 500
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

func configLokiBulkSize() (i int) {
	return intEnvOrDefault(EnvLokiBulkSize, DefaultLokiBulkSize)
}

func configLokiFlushIntervalMs() (i int) {
	return intEnvOrDefault(EnvLokiFlushIntervalMs, DefaultLokiFlushIntervalMs)
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
