package clients

import (
	"github.com/Gameye/gameye-sdk-go/models"
)

/**
 * Fetch the game state
 */
func (client GameyeClient) QueryGame() (
	err error,
	state *models.GameQueryState,
) {
	var anyState map[string]interface{}
	anyState, err = client.query(
		"game",
		nil,
	)
	if err != nil {
		return
	}

	state, err = models.CreateGameQueryState(&anyState)
	if err != nil {
		return
	}

	return
}

/**
 * Subscribe to the game state
 */
func (client GameyeClient) SubscribeGame() (
	err error,
	subscription *GameQuerySubscription,
) {
	var qs QuerySubscription
	qs, err = client.subscribe(
		"game",
		nil,
	)
	if err != nil {
		return
	}

	subscription = &GameQuerySubscription{
		qs,
	}

	return
}

type GameQuerySubscription struct {
	qs QuerySubscription
}

func (s *GameQuerySubscription) Cancel() {
	s.qs.Cancel()
}

func (s *GameQuerySubscription) NextState() (
	state *models.GameQueryState,
	err error,
) {
	var anyState map[string]interface{}
	anyState, err = s.qs.NextState()
	if err != nil {
		return
	}
	state, err = models.CreateGameQueryState(&anyState)
	if err != nil {
		return
	}

	return
}
