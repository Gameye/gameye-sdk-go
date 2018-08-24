package models

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplateQueryState(t *testing.T) {
	var err error
	defer func() {
		assert.NoError(t, err)
	}()

	anyState := make(map[string]interface{})
	{
		reader := strings.NewReader(templateStateJSONMock)
		decoder := json.NewDecoder(reader)
		err = decoder.Decode(&anyState)
		if err != nil {
			return
		}
	}

	var state *TemplateQueryState
	state, err = CreateTemplateQueryState(&anyState)
	assert.Equal(t, templateStateMock, state)
}
