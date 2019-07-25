package loki

import (
	"testing"

	. "github.com/BaritoLog/go-boilerplate/testkit"
	"github.com/BaritoLog/go-boilerplate/timekit"
)

func TestConvertBytesToTimber_GenerateTimestamp(t *testing.T) {
	timekit.FreezeUTC("2018-06-06T12:12:12Z")
	defer timekit.Unfreeze()

	timber, err := ConvertBytesToTimber(sampleRawTimber())
	FatalIfError(t, err)
	FatalIf(t, timber.Timestamp() != "2018-06-06T12:12:12Z", "wrong timber.Timestamp()")
}

func TestConvertBytesToTimber_JsonParseError(t *testing.T) {
	_, err := ConvertBytesToTimber([]byte(`invalid_json`))
	FatalIfWrongError(t, err, string(JsonParseError))
}

func TestConvertBytesToTimber_MissingContext(t *testing.T) {
	_, err := ConvertBytesToTimber([]byte(`{"hello":"world"}`))
	FatalIfWrongError(t, err, string(MissingContextError))
}

func TestConvertBytesToTimber_InvalidContext(t *testing.T) {
	_, err := ConvertBytesToTimber([]byte(`{"_ctx":{}}`))
	FatalIfWrongError(t, err, "Invalid Context Error: es_index_prefix is missing")
}

func sampleRawTimber() []byte {
	return []byte(`{
		"location": "some-location",
		"message":"some-message",
		"_ctx": {
			"kafka_topic": "some_topic",
			"kafka_partition": 3,
			"kafka_replication_factor": 1,
			"es_index_prefix": "some-type",
			"es_document_type": "some-type",
			"app_max_tps": 10,
			"app_secret": "some-secret-1234"
		}
	}`)

}

func sampleRawTimberCollection() []byte {
	return []byte(`{
		"items": [
			{
				"location": "some-location-1",
				"message":"some-message-1"
			},
			{
				"location": "some-location-2",
				"message":"some-message-2"
			}
		],
		"_ctx": {
			"kafka_topic": "some_topic",
			"kafka_partition": 3,
			"kafka_replication_factor": 1,
			"es_index_prefix": "some-type",
			"es_document_type": "some-type",
			"app_max_tps": 10,
			"app_secret": "some-secret-1234"
		}
	}`)

}
