package selectors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectMatchList(t *testing.T) {
	matchList := SelectMatchList(matchStateMock)
	assert.Equal(t, 2, len(matchList))
	for _, matchItem := range matchList {
		switch matchItem.MatchKey {
		case "test-match-123":
		case "test-match-456":
		default:
			assert.Fail(t, matchItem.MatchKey)
		}
	}
}

func TestSelectMatchListForGame(t *testing.T) {
	matchList := SelectMatchListForGame(matchStateMock, "test")
	assert.Equal(t, 1, len(matchList))
	for _, matchItem := range matchList {
		switch matchItem.MatchKey {
		case "test-match-123":
		default:
			assert.Fail(t, matchItem.MatchKey)
		}
	}
}

func TestSelectMatchItem(t *testing.T) {
	matchItem := SelectMatchItem(matchStateMock, "test-match-123")
	assert.NotNil(t, matchItem)
	if matchItem != nil {
		assert.Equal(t, "test-match-123", matchItem.MatchKey)
	}
}
