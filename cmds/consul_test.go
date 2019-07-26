package cmds

import (
	"net/http"
	"testing"

	. "github.com/BaritoLog/go-boilerplate/testkit"
)

func TestConsulLoki(t *testing.T) {
	ts := NewTestServer(http.StatusOK, []byte(`[
		{
			"ServiceAddress": "172.17.0.3",
			"ServicePort": 5000,
			"ServiceMeta": {
	        "http_schema": "https"
	    }
		}
	]`))
	defer ts.Close()

	url, err := consulLokiUrl(ts.URL, "name")
	FatalIfError(t, err)
	FatalIf(t, url != "https://172.17.0.3:5000", "wrong url")
}

func TestConsulLoki_NoHttpSchema(t *testing.T) {
	ts := NewTestServer(http.StatusOK, []byte(`[
		{
			"ServiceAddress": "172.17.0.3",
			"ServicePort": 5000
		}
	]`))
	defer ts.Close()

	url, err := consulLokiUrl(ts.URL, "name")
	FatalIfError(t, err)
	FatalIf(t, url != "http://172.17.0.3:5000", "wrong url")
}

func TestConsulLoki_NoService(t *testing.T) {
	ts := NewTestServer(http.StatusOK, []byte(`[]`))
	defer ts.Close()

	_, err := consulLokiUrl(ts.URL, "name")
	FatalIfWrongError(t, err, "No Service")
}
