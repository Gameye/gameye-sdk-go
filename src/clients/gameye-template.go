package clients

import (
	"github.com/Gameye/gameye-sdk-go/src/models"
)

// type TemplateQueryArg struct {
// 	GameKey string
// }

/**
 * Fetch template state
 * @param gameKey identifier of the game
 */
func (client GameyeClient) QueryTemplate(
	gameKey string,
) (
	err error,
	state *models.TemplateQueryState,
) {
	err = client.query(
		"template",
		map[string]string{
			"gameKey": gameKey,
		},
		state,
	)
	return
}

/**
 * Subscribe to template state
 * @param gameKey identifier of the game
 */
func (client GameyeClient) SubscribeTemplate(
	cancelChannel <-chan struct{},
	gameKey string,
) (
	err error,
	stateChannel chan<- *models.TemplateQueryState,
) {
	var state *models.TemplateQueryState
	var anyStateChannel <-chan interface{}

	anyStateChannel, err = client.subscribe(
		"template",
		map[string]string{
			"gameKey": gameKey,
		},
		state,
		cancelChannel,
	)

	go func() {
		for anyState := range anyStateChannel {
			stateChannel <- anyState.(*models.TemplateQueryState)
		}
	}()

	return
}
