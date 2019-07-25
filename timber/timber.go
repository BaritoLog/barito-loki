package timber

import (
	"fmt"
	"time"

	"github.com/BaritoLog/go-boilerplate/errkit"
)

const (
	InvalidContextError = errkit.Error("Invalid Context Error")
	MissingContextError = errkit.Error("Missing Context Error")
)

type Timber map[string]interface{}

type TimberCollection struct {
	Items  []Timber `json:"items"`
	Labels string   `json:"_labels"`
}

func (t Timber) SetTimestamp(timestamp string) {
	t["@timestamp"] = timestamp
}

func (t Timber) Timestamp() (s string) {
	s, _ = t["@timestamp"].(string)
	return
}

func (t Timber) Labels() (s string) {
	s, _ = t["_labels"].(string)
	return
}

func (t Timber) SetAppNameLabel(s string) {
	labels := fmt.Sprintf("{app_name=\"%s-%s\"}", s, time.Now().Format("2006.01.02"))
	t["_labels"] = labels
}

func (t Timber) InitAppNameLabel() (err error) {
	ctxMap, ok := t["_ctx"].(map[string]interface{})
	if !ok {
		err = MissingContextError
		return
	}

	esIndexPrefix, ok := ctxMap["es_index_prefix"].(string)
	if !ok {
		err = fmt.Errorf("es_index_prefix is missing")
		err = errkit.Concat(InvalidContextError, err)
		return
	}

	t.SetAppNameLabel(esIndexPrefix)
	delete(t, "_ctx")
	return
}
