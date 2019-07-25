package loki

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	pb "github.com/vwidjaya/barito-loki/timberproto"
	"github.com/BaritoLog/go-boilerplate/errkit"
	"github.com/golang/protobuf/ptypes"
)

const (
	JsonParseError = errkit.Error("JSON Parse Error")
)

func ConvertBytesToTimber(data []byte) (timber Timber, err error) {
	err = json.Unmarshal(data, &timber)
	if err != nil {
		err = errkit.Concat(JsonParseError, err)
		return
	}

	err = timber.InitAppNameLabel()
	if err != nil {
		return
	}

	if timber.Timestamp() == "" {
		timber.SetTimestamp(time.Now().UTC().Format(time.RFC3339))
	}

	return
}

func ConvertRequestToTimber(req *http.Request) (Timber, error) {
	body, _ := ioutil.ReadAll(req.Body)
	return ConvertBytesToTimber(body)
}

func ConvertTimberToLokiProto(timber Timber) *pb.Entry {
	timberMap := make(map[string]interface{})
	for k, v := range timber {
		timberMap[k] = v
	}

	delete(timberMap, "_labels")
	line, _ := json.Marshal(timberMap)

	return &pb.Entry{
		Timestamp: ptypes.TimestampNow(),
		Line:      string(line),
	}
}
