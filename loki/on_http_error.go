package loki

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.ErrorLevel)
}

type ForwardResult struct {
	Labels string `json:"labels"`
}

func onBadRequest(rw http.ResponseWriter, err error) {
	rw.WriteHeader(http.StatusBadRequest)
	rw.Write([]byte(err.Error()))
	log.Warn(err)
}

func onStoreError(rw http.ResponseWriter, err error) {
	rw.WriteHeader(http.StatusBadGateway)
	rw.Write([]byte(err.Error()))
	log.Warn(err)
}

func onSuccess(rw http.ResponseWriter, result ForwardResult) {
	rw.WriteHeader(http.StatusOK)

	b, _ := json.Marshal(result)
	rw.Write(b)
}
