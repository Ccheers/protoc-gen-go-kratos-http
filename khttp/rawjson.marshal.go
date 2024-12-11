package khttp

import "encoding/json"

func (x *RawJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(json.RawMessage(x.Json))
}
