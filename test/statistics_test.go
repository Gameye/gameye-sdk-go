package test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"../pkg/client/statistics"
	"./testutils"
)

func readStats() (result map[string]interface{}) {
	result, err := testutils.ReadFileAsJSON("./content/stats.json")
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func TestSelectsAllPlayers(test *testing.T) {
	state := statistics.StateWithStatistics(readStats())
	filtered := statistics.SelectPlayerList(state)

	assert.Equal(test, 10, len(filtered))
	for _, v := range filtered {
		if v.PlayerKey == "11" {
			assert.Equal(test, "Ivan", v.Name)
			return
		}
	}
}

func TestSelectsOnlyRequestedPlayers(test *testing.T) {
	state := statistics.StateWithStatistics(readStats())
	filtered, err := statistics.SelectPlayerListForTeam(state, "1")

	assert.Nil(test, err)
	assert.Equal(test, 5, len(filtered))
	assertions := 0
	for _, v := range filtered {
		if v.PlayerKey == "7" {
			assertions++
		}

		if v.Name == "Seth" {
			assertions++
		}

		if v.Name == "Zane" {
			assertions++
			test.Fail()
		}
		if v.PlayerKey == "8" {
			assertions++
			test.Fail()
		}
	}

	assert.Equal(test, 2, assertions)

	found, err := statistics.SelectPlayer(state, "9")
	assert.Nil(test, err)
	assert.Equal(test, "Vladimir", found.Name)
}

func TestSelectsAllTeams(test *testing.T) {
	state := statistics.StateWithStatistics(readStats())
	filtered := statistics.SelectTeamList(state)

	assert.Equal(test, 2, len(filtered))
	for _, v := range filtered {
		if v.TeamKey == "2" {
			assert.Equal(test, "Terrorists", v.Name)
			return
		}
	}
}
func TestSelectsOnlyRequestedTeams(test *testing.T) {
	state := statistics.StateWithStatistics(readStats())
	team, err := statistics.SelectTeam(state, "1")

	assert.Nil(test, err)
	assert.NotEmpty(test, team)
	assert.Equal(test, "1", team.TeamKey)
	assert.Equal(test, "Counter Terrorists", team.Name)
}

func TestSelectsRounds(test *testing.T) {
	state := statistics.StateWithStatistics(readStats())
	rounds := statistics.SelectRounds(state)

	assert.NotEmpty(test, rounds)
	assert.Equal(test, 2, rounds)
}

func TestSelectsRawStatistics(test *testing.T) {
	state := statistics.StateWithStatistics(readStats())
	raw, _ := statistics.SelectRawStatistics(state)

	unmarshaled := make(map[string]interface{})
	data := readStats()

	err := json.Unmarshal([]byte(raw), &unmarshaled)

	assert.Nil(test, err)
	assert.Equal(test, data, unmarshaled)
}
