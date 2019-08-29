package loki

import (
	"bytes"
	"net/http"
	"strings"
	"testing"
	"time"

	. "github.com/BaritoLog/go-boilerplate/testkit"
	"github.com/BaritoLog/go-boilerplate/timekit"
	"github.com/golang/mock/gomock"
	"github.com/grafana/loki/pkg/promtail/api"
	promtail "github.com/grafana/loki/pkg/promtail/client/fake"
	"github.com/prometheus/common/model"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.ErrorLevel)
}

func TestBaritoLokiService_ServeHTTP_OnBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := &baritoLokiService{
		ptClient: FakeLokiClient(),
	}
	defer service.Close()

	req, _ := http.NewRequest("POST", "/", strings.NewReader(`invalid-body`))
	resp := RecordResponse(service.ServeHTTP, req)

	FatalIfWrongResponseStatus(t, resp, http.StatusBadRequest)
}

func TestBaritoLokiService_ServeHTTP_OnStoreError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := &baritoLokiService{}
	defer service.Close()

	req, _ := http.NewRequest("POST", "/", bytes.NewReader(sampleRawTimber()))
	resp := RecordResponse(service.ServeHTTP, req)

	FatalIfWrongResponseStatus(t, resp, http.StatusBadGateway)
}

func TestBaritoLokiService_ServeHTTP_OnSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := &baritoLokiService{
		ptClient: FakeLokiClient(),
	}
	defer service.Close()

	req, _ := http.NewRequest("POST", "/", bytes.NewReader(sampleRawTimber()))
	resp := RecordResponse(service.ServeHTTP, req)

	FatalIfWrongResponseStatus(t, resp, http.StatusOK)
}

func TestBaritoLokiService_ServeHTTP_ProduceBatch_OnSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := &baritoLokiService{
		ptClient: FakeLokiClient(),
	}
	defer service.Close()

	req, _ := http.NewRequest("POST", "/produce_batch", bytes.NewReader(sampleRawTimberCollection()))
	resp := RecordResponse(service.ServeHTTP, req)

	FatalIfWrongResponseStatus(t, resp, http.StatusOK)
}

func TestBaritoLokiService_Start(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := &baritoLokiService{
		addr:     ":24400",
		ptClient: FakeLokiClient(),
	}
	defer service.Close()

	var err error
	go func() {
		err = service.Start()
	}()
	defer service.Close()

	FatalIfError(t, err)

	timekit.Sleep("1ms")
	resp, err := http.Get("http://:24400")
	FatalIfError(t, err)
	FatalIfWrongResponseStatus(t, resp, http.StatusBadRequest)
}

func FakeLokiClient() *promtail.Client {
	onHandleFunc := api.EntryHandlerFunc(func(labels model.LabelSet, time time.Time, entry string) error { return nil })
	onStopFunc := func() {}
	return &promtail.Client{
		OnHandleEntry: onHandleFunc,
		OnStop:        onStopFunc,
	}
}
