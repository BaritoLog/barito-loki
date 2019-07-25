package loki

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/snappy"

	pb "github.com/BaritoLog/barito-loki/timberproto"
	log "github.com/sirupsen/logrus"
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

type Loki interface {
	Store(timber Timber)
}

type lokiClient struct {
	config  *lokiConfig
	entries chan *lokiEntry
	client  *http.Client
}

func NewLoki(conf lokiConfig) Loki {
	lkClient := lokiClient{
		config:  &conf,
		entries: make(chan *lokiEntry, conf.bulkSize),
		client:  &http.Client{},
	}

	go lkClient.run()

	return &lkClient
}

type lokiEntry struct {
	labels string
	entry  *pb.Entry
}

func (c *lokiClient) Store(timber Timber) {
	labels := timber.Labels()
	entry := ConvertTimberToLokiProto(timber)

	c.entries <- &lokiEntry{
		labels: labels,
		entry:  entry,
	}
}

func (c *lokiClient) run() {
	batch := map[string]*pb.Stream{}
	batchSize := 0
	maxWait := time.NewTimer(c.config.flushInterval)

	defer func() {
		if batchSize > 0 {
			c.send(batch)
		}
	}()

	for {
		select {
		case e := <-c.entries:
			stream, ok := batch[e.labels]
			if !ok {
				stream = &pb.Stream{
					Labels: e.labels,
				}
				batch[e.labels] = stream
			}

			stream.Entries = append(stream.Entries, e.entry)
			batchSize++

			if batchSize >= c.config.bulkSize {
				c.send(batch)
				batch = map[string]*pb.Stream{}
				batchSize = 0
				maxWait.Reset(c.config.flushInterval)
			}
		case <-maxWait.C:
			if batchSize > 0 {
				c.send(batch)
				batch = map[string]*pb.Stream{}
				batchSize = 0
			}
			maxWait.Reset(c.config.flushInterval)
		}
	}
}

func (c *lokiClient) send(batch map[string]*pb.Stream) {
	buf, err := encodeBatch(batch)
	if err != nil {
		log.Warnf("Loki Client - unable to marshal: %s\n", err)
		return
	}

	resp, body, err := c.sendReq(buf)
	if err != nil {
		log.Warnf("Loki Client - unable to send an HTTP request: %s\n", err)
		return
	}

	if resp.StatusCode != 204 {
		log.Warnf("Loki Client - unexpected HTTP status code: %d, message: %s\n", resp.StatusCode, body)
		return
	}
}

func encodeBatch(batch map[string]*pb.Stream) ([]byte, error) {
	req := pb.PushRequest{
		Streams: make([]*pb.Stream, 0, len(batch)),
	}

	for _, stream := range batch {
		req.Streams = append(req.Streams, stream)
	}

	buf, err := proto.Marshal(&req)
	if err != nil {
		return nil, err
	}

	buf = snappy.Encode(nil, buf)
	return buf, nil
}

func (c *lokiClient) sendReq(buf []byte) (resp *http.Response, resBody []byte, err error) {
	resp, err = http.Post(c.config.pushURL, ContentType, bytes.NewBuffer(buf))
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	resBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return resp, resBody, nil
}
