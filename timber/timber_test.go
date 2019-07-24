package timber

import (
	"testing"

	. "github.com/BaritoLog/go-boilerplate/testkit"
)

func TestInitAppNameLabel_ESIndexPrefixIsMissing(t *testing.T) {
	timber := sampleTimber()
	ctx := timber["_ctx"].(map[string]interface{})
	delete(ctx, "es_index_prefix")

	err := timber.InitAppNameLabel()
	FatalIfWrongError(t, err, string("Invalid Context Error: es_index_prefix is missing"))
}

func TestInitAppNameLabel_ValidESIndexPrefix(t *testing.T) {
	timber := sampleTimber()
	err := timber.InitAppNameLabel()
	FatalIfError(t, err)
}

func sampleTimber() (timber Timber) {
	timber = make(map[string]interface{})
	timber["location"] = "some-location"
	timber["message"] = "some-message"
	timber["_ctx"] = map[string]interface{}{
		"es_index_prefix": "some-prefix",
	}
	return
}
