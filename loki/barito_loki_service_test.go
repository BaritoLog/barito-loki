package loki

import (
	"bytes"
	"net/http"
	"strings"
	"testing"

	. "github.com/BaritoLog/go-boilerplate/testkit"
	"github.com/BaritoLog/go-boilerplate/timekit"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.ErrorLevel)
}

func TestBaritoLokiService_ServeHTTP_OnBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := &baritoLokiService{}
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

	lkConfig := NewLokiConfig("http://localhost:3100", 500, 500)
	service := &baritoLokiService{
		lkClient: NewLoki(lkConfig),
	}
	defer service.Close()

	req, _ := http.NewRequest("POST", "/", bytes.NewReader(sampleRawTimber()))
	resp := RecordResponse(service.ServeHTTP, req)

	FatalIfWrongResponseStatus(t, resp, http.StatusOK)
}

func TestBaritoLokiService_Start(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := &baritoLokiService{
		addr: ":24400",
	}

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
