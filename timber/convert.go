package timber

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/BaritoLog/go-boilerplate/errkit"
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
