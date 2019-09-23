package patch

import "encoding/json"

type Patch struct {
	Path  []interface{}    `json:"path"`
	Value *json.RawMessage `json:"value"`
}
