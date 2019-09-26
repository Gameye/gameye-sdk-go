package event

import "encoding/json"

// UnionEvent is used for partial deserialization
type UnionEvent struct {
	Type    string           `json:"type"`
	Payload *json.RawMessage `json:"payload"`
}
