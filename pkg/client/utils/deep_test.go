package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetIn(t *testing.T) {
	initialState := map[string]interface{}{}

	actual := SetIn(initialState, []string{"1", "2"}, "yeah")
	assert.Equal(t, map[string]interface{}{}, initialState)

	expected := map[string]interface{}{
		"1": map[string]interface{}{
			"2": "yeah",
		},
	}
	assert.Equal(t, expected, actual)

	actual = SetIn(actual, []string{"1", "2"}, "yo")
	expected = map[string]interface{}{
		"1": map[string]interface{}{
			"2": "yo",
		},
	}
	assert.Equal(t, expected, actual)

	actual = SetIn(actual, []string{"3"}, "ya")
	expected = map[string]interface{}{
		"1": map[string]interface{}{
			"2": "yo",
		},
		"3": "ya",
	}
	assert.Equal(t, expected, actual)

}
