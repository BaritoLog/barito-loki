package loki

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/jsonpb"
	stpb "github.com/golang/protobuf/ptypes/struct"
	pb "github.com/vwidjaya/barito-proto/producer"
)

func GenerateLokiLabels(tCtx *pb.TimberContext) string {
	return fmt.Sprintf("{app_name=\"%s-%s\"}", tCtx.GetEsIndexPrefix(), time.Now().Format("2006.01.02"))
}

func SerializeTimberContents(timber *pb.Timber) string {
	m := jsonpb.Marshaler{}
	content := timber.GetContent()

	ts := &stpb.Value{
		Kind: &stpb.Value_StringValue{
			StringValue: time.Now().UTC().Format(time.RFC3339),
		},
	}
	content.Fields["@timestamp"] = ts

	line, _ := m.MarshalToString(content)
	return line
}
