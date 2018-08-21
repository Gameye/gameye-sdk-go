package clients

import (
	"github.com/Gameye/gameye-sdk-go/src/models"
)

/**
 * Fetch the game state
 */
func (client GameyeClient) QueryGame() (
	err error,
	state *models.GameQueryState,
) {
	err = client.query(
		"game",
		map[string]string{},
		state,
	)
	return
}

/**
 * Subscribe to the game state
 */
func (client GameyeClient) SubscribeGame() (
	err error,
	state *models.GameQueryState,
) {
	err = client.subscribe(
		"game",
		map[string]string{},
		state,
	)
	return
}
