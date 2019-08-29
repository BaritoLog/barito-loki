package loki

import (
	"fmt"
	"strings"
	"testing"
	"time"

	. "github.com/BaritoLog/go-boilerplate/testkit"
	"github.com/grafana/loki/pkg/promtail/api"
	promtail "github.com/grafana/loki/pkg/promtail/client/fake"
	"github.com/prometheus/common/model"
	pb "github.com/vwidjaya/barito-proto/producer"
)

func TestBaritoLokiService_Produce_OnStoreError(t *testing.T) {
	srv := &baritoLokiService{}

	_, err := srv.Produce(nil, pb.SampleTimberProto())
	FatalIfWrongGrpcError(t, onStoreErrorGrpc(fmt.Errorf("")), err)
}

func TestBaritoLokiService_Produce_OnSuccess(t *testing.T) {
	srv := &baritoLokiService{
		ptClient: FakePromtailClient(),
	}

	resp, err := srv.Produce(nil, pb.SampleTimberProto())
	FatalIfError(t, err)

	expected := GenerateLokiLabels(pb.SampleTimberProto().GetContext())
	FatalIf(t, resp.GetTopic() != expected, "wrong result.Topic (Loki labels)")
}

func TestBaritoLokiService_ProduceBatch_OnStoreError(t *testing.T) {
	srv := &baritoLokiService{}

	_, err := srv.ProduceBatch(nil, pb.SampleTimberCollectionProto())
	FatalIfWrongGrpcError(t, onStoreErrorGrpc(fmt.Errorf("")), err)
}

func TestBaritoLokiService_ProduceBatch_OnSuccess(t *testing.T) {
	srv := &baritoLokiService{
		ptClient: FakePromtailClient(),
	}

	resp, err := srv.ProduceBatch(nil, pb.SampleTimberCollectionProto())
	FatalIfError(t, err)

	expected := GenerateLokiLabels(pb.SampleTimberCollectionProto().GetContext())
	FatalIf(t, resp.GetTopic() != expected, "wrong result.Topic (Loki labels)")
}

func TestBaritoLokiService_Start(t *testing.T) {
	service := &baritoLokiService{
		grpcAddr: ":24400",
		ptClient: FakePromtailClient(),
	}
	defer service.Close()

	var err error
	go func() {
		err = service.Start()
	}()
	defer service.Close()

	FatalIfError(t, err)
}

func FakePromtailClient() *promtail.Client {
	onHandleFunc := api.EntryHandlerFunc(func(labels model.LabelSet, time time.Time, entry string) error { return nil })
	onStopFunc := func() {}
	return &promtail.Client{
		OnHandleEntry: onHandleFunc,
		OnStop:        onStopFunc,
	}
}

func FatalIfWrongGrpcError(t *testing.T, expected error, actual error) {
	expFields := strings.Fields(expected.Error())[:5]
	expStr := strings.Join(expFields, " ")

	actFields := strings.Fields(actual.Error())[:5]
	actStr := strings.Join(actFields, " ")

	if expStr != actStr {
		t.Errorf("expected gRPC response code %v, received %v.", expFields[4], actFields[4])
	}
}
