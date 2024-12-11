package exampleapi

import (
	"encoding/json"
	"testing"

	"github.com/Ccheers/protoc-gen-go-kratos-http/khttp"
)

func TestMarshal(t *testing.T) {

	got, _ := json.Marshal(&HelloWorldRequest{
		Name: "test",
		Raw:  &khttp.RawJson{Json: []byte(`{"key":"value"}`)},
	})
	need := `{"name":"test","raw":{"key":"value"}}`
	if string(got) != need {
		t.Errorf("need=%s, gos=%s", need, string(got))
	}
	t.Log(string(got))

	got, _ = json.Marshal(&HelloWorldRequest{
		Name: "test",
		Raw:  nil,
	})
	need = `{"name":"test"}`
	if string(got) != need {
		t.Errorf("need=%s, gos=%s", need, string(got))
	}

	t.Log(string(got))

}
