package models

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatisticQueryState(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	anyState := make(map[string]interface{})
	{
		reader := strings.NewReader(statisticStateJSONMock)
		decoder := json.NewDecoder(reader)
		err = decoder.Decode(&anyState)
		if err != nil {
			return
		}
	}

	var state *StatisticQueryState
	state, err = CreateStatisticQueryState(&anyState)

	assert.Equal(t, statisticStateMock, state)
}
