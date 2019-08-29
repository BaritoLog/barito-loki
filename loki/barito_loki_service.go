package loki

import (
	"net/http"
	"net/url"
	"time"

	"github.com/BaritoLog/go-boilerplate/errkit"
	"github.com/cortexproject/cortex/pkg/util"
	"github.com/cortexproject/cortex/pkg/util/flagext"
	"github.com/prometheus/common/model"

	logkit "github.com/go-kit/kit/log/logrus"
	promtail "github.com/grafana/loki/pkg/promtail/client"
	logrus "github.com/sirupsen/logrus"
)

const (
	ErrLokiClient = errkit.Error("Loki Client Failed")
)

type BaritoLokiService interface {
	Start() error
	Close()
	ServeHTTP(rw http.ResponseWriter, req *http.Request)
}

type baritoLokiService struct {
	addr     string
	server   *http.Server
	ptConfig promtail.Config
	ptClient promtail.Client
}

func NewBaritoLokiService(params map[string]interface{}) (srv BaritoLokiService, err error) {
	ptConfig, err := parseLokiConfig(params)
	if err != nil {
		return
	}

	srv = &baritoLokiService{
		addr:     params["serviceAddr"].(string),
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

	url, err := url.Parse(lokiUrl)
	if err != nil {
		return
	}

	promtailURL := flagext.URLValue{
		URL: url,
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

func (s *baritoLokiService) Start() (err error) {
	err = s.initPromtailClient()
	if err != nil {
		return
	}

	server := s.initHttpServer()
	return server.ListenAndServe()
}

func (s *baritoLokiService) Close() {
	if s.ptClient != nil {
		s.ptClient.Stop()
	}

	if s.server != nil {
		s.server.Close()
	}
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

func (s *baritoLokiService) initHttpServer() (server *http.Server) {
	server = &http.Server{
		Addr:    s.addr,
		Handler: s,
	}

	s.server = server
	return
}

func (s *baritoLokiService) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if s.ptClient == nil {
		onStoreError(rw, ErrLokiClient)
		return
	}

	var labels string

	if req.URL.Path == "/produce_batch" {
		timberCollection, err := ConvertBatchRequestToTimberCollection(req)
		if err != nil {
			onBadRequest(rw, err)
			return
		}

		esIndexPrefix := timberCollection.Context["es_index_prefix"].(string)
		labels = generateLabelForTimber(esIndexPrefix)

		var ls model.LabelSet
		ls.UnmarshalJSON([]byte(labels))

		for _, timber := range timberCollection.Items {
			timber.SetAppNameLabel(labels)
			if timber.Timestamp() == "" {
				timber.SetTimestamp(time.Now().UTC().Format(time.RFC3339))
			}

			ts := time.Now().UTC()
			line := ConvertTimberToLokiEntryLine(timber)

			_ = s.ptClient.Handle(ls, ts, line)
		}
	} else {
		timber, err := ConvertRequestToTimber(req)
		if err != nil {
			onBadRequest(rw, err)
			return
		}

		var ls model.LabelSet
		labels = timber.Labels()
		ls.UnmarshalJSON([]byte(labels))

		ts := time.Now().UTC()
		line := ConvertTimberToLokiEntryLine(timber)

		_ = s.ptClient.Handle(ls, ts, line)
	}

	onSuccess(rw, ForwardResult{
		Labels: labels,
	})
}
