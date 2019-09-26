package patch

import "encoding/json"

// Patch represents a json patch at a path
type Patch struct {
	Path  []interface{}    `json:"path"`
	Value *json.RawMessage `json:"value"`
}
