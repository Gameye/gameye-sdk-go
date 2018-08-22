package selectors

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Gameye/gameye-sdk-go/src/models"
)

func TestSelectLocationListForGame(t *testing.T) {
	locationList := SelectLocationListForGame(&models.GameStateMock, "test")
	assert.Equal(t, 1, len(locationList))
	for _, locationItem := range locationList {
		switch locationItem.LocationKey {
		case "local":
		default:
			assert.Fail(t, locationItem.LocationKey)
		}
	}
}
