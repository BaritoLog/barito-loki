package loki

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/BaritoLog/go-boilerplate/errkit"
	"github.com/cortexproject/cortex/pkg/util"
	"github.com/cortexproject/cortex/pkg/util/flagext"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/prometheus/common/model"
	"google.golang.org/grpc"

	logkit "github.com/go-kit/kit/log/logrus"
	promtail "github.com/grafana/loki/pkg/promtail/client"
	logrus "github.com/sirupsen/logrus"
	pb "github.com/vwidjaya/barito-proto/producer"
)

const (
	ErrInitGrpc     = errkit.Error("Failed to listen to gRPC address")
	ErrNilPromtail  = errkit.Error("Promtail not found")
	ErrRegisterGrpc = errkit.Error("Error registering gRPC server endpoint into reverse proxy")
	ErrReverseProxy = errkit.Error("Error serving REST reverse proxy")
)

type BaritoLokiService interface {
	pb.ProducerServer
	Start() error
	LaunchREST() error
	Close()
}

type baritoLokiService struct {
	grpcAddr string
	restAddr string

	grpcServer   *grpc.Server
	reverseProxy *http.Server

	ptConfig promtail.Config
	ptClient promtail.Client
}

func NewBaritoLokiService(params map[string]interface{}) (srv BaritoLokiService, err error) {
	ptConfig, err := parseLokiConfig(params)
	if err != nil {
		return
	}

	srv = &baritoLokiService{
		grpcAddr: params["grpcAddr"].(string),
		restAddr: params["restAddr"].(string),
		ptConfig: ptConfig,
	}

	return
}

func parseLokiConfig(params map[string]interface{}) (cfg promtail.Config, err error) {
	lokiUrl := params["lokiUrl"].(string)
	batchWaitMs := params["batchWaitMs"].(int)
	batchSize := params["batchSize"].(int)
	minBackoffMs := params["minBackoffMs"].(int)
	maxBackoffMs := params["maxBackoffMs"].(int)
	maxRetries := params["maxRetries"].(int)
	timeoutMs := params["timeoutMs"].(int)

	var promtailURL flagext.URLValue
	err = promtailURL.Set(fmt.Sprintf("%s/api/prom/push", lokiUrl))
	if err != nil {
		return
	}

	cfg = promtail.Config{
		URL:       promtailURL,
		BatchWait: time.Duration(batchWaitMs) * time.Millisecond,
		BatchSize: batchSize,
		BackoffConfig: util.BackoffConfig{
			MinBackoff: time.Duration(minBackoffMs) * time.Millisecond,
			MaxBackoff: time.Duration(maxBackoffMs) * time.Millisecond,
			MaxRetries: maxRetries,
		},
		Timeout: time.Duration(timeoutMs) * time.Millisecond,
	}

	return
}

func (s *baritoLokiService) initPromtailClient() (err error) {
	if s.ptClient != nil {
		return nil
	}

	logger := logkit.NewLogrusLogger(logrus.New())

	ptClient, err := promtail.New(s.ptConfig, logger)
	if err != nil {
		return
	}

	s.ptClient = ptClient
	return
}

func (s *baritoLokiService) initGrpcServer() (lis net.Listener, srv *grpc.Server, err error) {
	lis, err = net.Listen("tcp", s.grpcAddr)
	if err != nil {
		return
	}

	srv = grpc.NewServer()
	pb.RegisterProducerServer(srv, s)

	s.grpcServer = srv
	return
}

func (s *baritoLokiService) Start() (err error) {
	err = s.initPromtailClient()
	if err != nil {
		return
	}

	lis, grpcSrv, err := s.initGrpcServer()
	if err != nil {
		err = errkit.Concat(ErrInitGrpc, err)
		return
	}

	return grpcSrv.Serve(lis)
}

func (s *baritoLokiService) LaunchREST() (err error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = pb.RegisterProducerHandlerFromEndpoint(ctx, mux, "localhost"+s.grpcAddr, opts)
	if err != nil {
		err = errkit.Concat(ErrRegisterGrpc, err)
		return
	}

	s.reverseProxy = &http.Server{
		Addr:    s.restAddr,
		Handler: mux,
	}

	err = s.reverseProxy.ListenAndServe()
	if err != nil {
		err = errkit.Concat(ErrReverseProxy, err)
	}
	return
}

func (s *baritoLokiService) Close() {
	if s.reverseProxy != nil {
		s.reverseProxy.Close()
	}

	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}

	if s.ptClient != nil {
		s.ptClient.Stop()
	}
}

func (s *baritoLokiService) Produce(_ context.Context, timber *pb.Timber) (resp *pb.ProduceResult, err error) {
	if s.ptClient == nil {
		err = onStoreErrorGrpc(ErrNilPromtail)
		return
	}

	var ls model.LabelSet
	labels := GenerateLokiLabels(timber.GetContext())
	ls.UnmarshalJSON([]byte(labels))

	ts := time.Now().UTC()
	line := SerializeTimberContents(timber)

	_ = s.ptClient.Handle(ls, ts, line)

	resp = &pb.ProduceResult{
		Topic: labels,
	}
	return
}

func (s *baritoLokiService) ProduceBatch(_ context.Context, timberCollection *pb.TimberCollection) (resp *pb.ProduceResult, err error) {
	if s.ptClient == nil {
		err = onStoreErrorGrpc(ErrNilPromtail)
		return
	}

	var ls model.LabelSet
	labels := GenerateLokiLabels(timberCollection.GetContext())
	ls.UnmarshalJSON([]byte(labels))

	for _, timber := range timberCollection.GetItems() {
		ts := time.Now().UTC()
		line := SerializeTimberContents(timber)

		_ = s.ptClient.Handle(ls, ts, line)
	}

	resp = &pb.ProduceResult{
		Topic: labels,
	}
	return
}
