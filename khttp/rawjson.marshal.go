package khttp

import "encoding/json"

func (x *RawJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(json.RawMessage(x.Json))
}

func (x *RawJson) UnmarshalJSON(bs []byte) error {
	if x == nil {
		return nil
	}
	x.Json = bs
	return nil
}

func NewRawJSON(message json.RawMessage) *RawJson {
	return &RawJson{Json: message}
}
