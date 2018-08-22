package models

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestMatchQueryState(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	anyState := make(map[string]interface{})
	{
		reader := strings.NewReader(MatchStateJSONMock)
		decoder := json.NewDecoder(reader)
		err = decoder.Decode(&anyState)
		if err != nil {
			return
		}
	}

	var state MatchQueryState
	mapstructure.Decode(&anyState, &state)

	assert.Equal(t, MatchStateMock, state)
}
