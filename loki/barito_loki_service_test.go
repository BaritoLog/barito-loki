package loki

import (
	"bytes"
	"net/http"
	"strings"
	"testing"

	. "github.com/BaritoLog/go-boilerplate/testkit"
	"github.com/golang/mock/gomock"
)

func TestBaritoLokiService_ServeHTTP_OnBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := &baritoLokiService{}
	defer service.Close()

	req, _ := http.NewRequest("POST", "/", strings.NewReader(`invalid-body`))
	resp := RecordResponse(service.ServeHTTP, req)

	FatalIfWrongResponseStatus(t, resp, http.StatusBadRequest)
}

func TestBaritoLokiService_ServeHTTP_OnSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := &baritoLokiService{}
	defer service.Close()

	req, _ := http.NewRequest("POST", "/", bytes.NewReader(sampleRawTimber()))
	resp := RecordResponse(service.ServeHTTP, req)

	FatalIfWrongResponseStatus(t, resp, http.StatusOK)
}

func sampleRawTimber() []byte {
	return []byte(`{
		"location": "some-location",
		"message":"some-message",
		"_ctx": {
			"kafka_topic": "some_topic",
			"kafka_partition": 3,
			"kafka_replication_factor": 1,
			"es_index_prefix": "some-type",
			"es_document_type": "some-type",
			"app_max_tps": 10,
			"app_secret": "some-secret-1234"
		}
	}`)

}
