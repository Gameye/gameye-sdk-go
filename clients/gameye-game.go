package clients

import (
	"github.com/Gameye/gameye-sdk-go/models"
)

/*
QueryGame will Fetch the game state
*/
func (client GameyeClient) QueryGame() (
	state *models.GameQueryState,
	err error,
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

/*
SubscribeGame subscribes to the game state
*/
func (client GameyeClient) SubscribeGame() (
	subscription *GameQuerySubscription,
	err error,
) {
	var qs *querySubscription
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

/*
GameQuerySubscription is a subscription to the game state
*/
type GameQuerySubscription struct {
	qs *querySubscription
}

/*
Cancel will end and cleanup the subscription
*/
func (s *GameQuerySubscription) Cancel() {
	s.qs.Cancel()
}

/*
NextState will return the next state
*/
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
