package models

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchQueryState(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	anyState := make(map[string]interface{})
	{
		reader := strings.NewReader(matchStateJSONMock)
		decoder := json.NewDecoder(reader)
		err = decoder.Decode(&anyState)
		if err != nil {
			return
		}
	}

	var state *MatchQueryState
	state, err = CreateMatchQueryState(&anyState)

	assert.Equal(t, matchStateMock, state)
}
