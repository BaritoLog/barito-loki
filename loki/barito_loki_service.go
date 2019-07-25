package loki

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli"
)

func Start(c *cli.Context) (err error) {
	fmt.Println("Loki client started.")
	return
}

type BaritoLokiService interface {
	Start() error
	Close()
	ServeHTTP(rw http.ResponseWriter, req *http.Request)
}

type baritoLokiService struct {
	addr   string
	server *http.Server
}

func NewBaritoLokiService(addr string) BaritoLokiService {
	return &baritoLokiService{
		addr: addr,
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

	onSuccess(rw, ForwardResult{
		Labels: timber.Labels(),
	})
}
