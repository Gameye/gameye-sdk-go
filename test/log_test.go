package test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"../pkg/client/logs"
	"./testutils"
)

func readLogs() (json map[string]interface{}) {
	json, err := testutils.ReadFileAsJSON("./content/logs.json")
	if err != nil {
		log.Fatal(err)
	}
	return json
}

func TestSelectsAllLogs(test *testing.T) {
	state := logs.StateWithLogs(readLogs())
	filtered := logs.SelectAllLogs(state)

	assert.Equal(test, 1095, len(filtered))
	assert.Equal(test, 561, filtered[560].LineKey)
	assert.Equal(test, "$L 09/16/2019 - 12:39:05: \"Joe<4><BOT><TERRORIST>\" dropped \"vesthelm\"", filtered[896].Payload)
}

func TestSelectsOnlyRequestedLogs(test *testing.T) {
	state := logs.StateWithLogs(readLogs())
	filtered := logs.SelectLogsSince(state, 912)

	assert.Equal(test, 1095 - 912, len(filtered))
}