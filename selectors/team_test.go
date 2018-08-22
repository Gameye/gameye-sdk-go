package selectors

import (
	"testing"

	"github.com/Gameye/gameye-sdk-go/models"
	"github.com/stretchr/testify/assert"
)

func TestSelectTeamList(t *testing.T) {
	teamList := SelectTeamList(&models.StatisticStateMock)
	assert.Equal(t, 2, len(teamList))
	for _, teamItem := range teamList {
		switch teamItem.TeamKey {
		case "1":
			assert.Equal(t, "TeamA", teamItem.Name)
		case "2":
			assert.Equal(t, "TeamB", teamItem.Name)
		default:
			assert.Fail(t, teamItem.TeamKey)
		}
	}
}

func TestSelectTeamItem(t *testing.T) {
	teamItem := SelectTeamItem(&models.StatisticStateMock, "2")
	assert.NotNil(t, teamItem)
	if teamItem != nil {
		assert.Equal(t, "2", teamItem.TeamKey)
		assert.Equal(t, "TeamB", teamItem.Name)
	}
}
