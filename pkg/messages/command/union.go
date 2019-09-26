package command

import "encoding/json"

// UnionCommand is used for partial deserialization
type UnionCommand struct {
	Type    string           `json:"type"`
	Payload *json.RawMessage `json:"payload"`
}
