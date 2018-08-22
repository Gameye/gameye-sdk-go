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
		nil,
		state,
	)
	return
}

/**
 * Subscribe to the game state
 */
func (client GameyeClient) SubscribeGame(
	cancelChannel <-chan struct{},
) (
	err error,
	stateChannel chan<- models.GameQueryState,
) {
	var state *models.GameQueryState
	var anyStateChannel <-chan interface{}

	anyStateChannel, err = client.subscribe(
		"game",
		nil,
		state,
		cancelChannel,
	)

	go func() {
		for anyState := range anyStateChannel {
			stateChannel <- anyState.(models.GameQueryState)
		}
	}()

	return
}
