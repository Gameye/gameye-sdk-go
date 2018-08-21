package clients

import (
	"github.com/Gameye/gameye-sdk-go/src/models"
	"github.com/Gameye/gameye-sdk-go/src/utils"
)

/**
 * Fetch the game state
 */
func (client GameyeClient) QueryGame() (
	err error,
	state *models.GameQueryState,
) {
	client.query("game", &utils.EmptyStruct)
	return
}

/**
 * Subscribe to the game state
 */
func (client GameyeClient) SubscribeGame() (
	err error,
	state *models.GameQueryState,
) {
	client.subscribe("game", &utils.EmptyStruct)
	return
}
