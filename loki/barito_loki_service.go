package loki

import (
	"fmt"
	"net/http"

	"github.com/BaritoLog/go-boilerplate/errkit"
	"github.com/BaritoLog/go-boilerplate/srvkit"
	"github.com/urfave/cli"
)

const (
	Address       = ":24400"
	ErrLokiClient = errkit.Error("Loki Client Failed")
)

func Start(c *cli.Context) (err error) {
	fmt.Println("Loki client started.")

	lkConfig := NewLokiConfig("http://localhost:3100", 500, 500)
	service := NewBaritoLokiService(Address, lkConfig)

	err = service.Start()
	if err != nil {
		return
	}
	srvkit.AsyncGracefulShutdown(service.Close)

	return
}

type BaritoLokiService interface {
	Start() error
	Close()
	ServeHTTP(rw http.ResponseWriter, req *http.Request)
}

type baritoLokiService struct {
	addr     string
	server   *http.Server
	lkClient Loki
}

func NewBaritoLokiService(addr string, lkConfig lokiConfig) BaritoLokiService {
	return &baritoLokiService{
		addr:     addr,
		lkClient: NewLoki(lkConfig),
	}
}

func (s *baritoLokiService) Start() (err error) {
	server := s.initHttpServer()
	return server.ListenAndServe()
}

func (a *baritoLokiService) Close() {
	if a.server != nil {
		a.server.Close()
	}
}

func (s *baritoLokiService) initHttpServer() (server *http.Server) {
	server = &http.Server{
		Addr:    s.addr,
		Handler: s,
	}

	s.server = server
	return
}

func (s *baritoLokiService) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	timber, err := ConvertRequestToTimber(req)
	if err != nil {
		onBadRequest(rw, err)
		return
	}

	if s.lkClient == nil {
		onStoreError(rw, ErrLokiClient)
		return
	}

	s.lkClient.Store(timber)

	onSuccess(rw, ForwardResult{
		Labels: timber.Labels(),
	})
}
