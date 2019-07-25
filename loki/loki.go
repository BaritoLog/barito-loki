package loki

import (
	"fmt"
	"net/http"
	"time"

	pb "github.com/BaritoLog/barito-loki/timberproto"
)

const (
	ContentType = "application/x-protobuf"
)

type lokiConfig struct {
	pushURL       string
	bulkSize      int
	flushInterval time.Duration
}

func NewLokiConfig(lkUrl string, bulkSize int, flushMs int) lokiConfig {
	return lokiConfig{
		pushURL:       fmt.Sprintf("%s/api/prom/push", lkUrl),
		bulkSize:      bulkSize,
		flushInterval: time.Duration(flushMs) * time.Millisecond,
	}
}

type lokiClient struct {
	config  *lokiConfig
	entries chan *lokiEntry
	client  *http.Client
}

func NewLoki(conf lokiConfig) (lkClient lokiClient) {
	lkClient = lokiClient{
		config:  &conf,
		entries: make(chan *lokiEntry, conf.bulkSize),
		client:  &http.Client{},
	}

	return
}

type lokiEntry struct {
	labels string
	entry  *pb.Entry
}
