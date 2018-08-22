package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetIn(test *testing.T) {
	state := map[string]interface{}{}
	value := "yeah"

	actual := SetIn(state, []string{"1", "2"}, value)
	assert.Equal(test, map[string]interface{}{}, state)

	expected := map[string]interface{}{
		"1": map[string]interface{}{
			"2": value,
		},
	}
	assert.Equal(test, expected, actual)
}
