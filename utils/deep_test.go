package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetIn(t *testing.T) {
	state := map[string]interface{}{}
	value := "yeah"

	actual := SetIn(state, []string{"1", "2"}, value)
	assert.Equal(t, map[string]interface{}{}, state)

	expected := map[string]interface{}{
		"1": map[string]interface{}{
			"2": value,
		},
	}
	assert.Equal(t, expected, actual)
}
