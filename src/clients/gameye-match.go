package clients

import (
	"github.com/Gameye/gameye-sdk-go/src/models"
)

type StartMatchPayload struct {
	MatchKey     string
	GameKey      string
	LocationKeys []string
	TemplateKey  string
	Config       map[string]interface{}
}

/**
 * Start a match
 * @param matchKey a unique identifier for this match, you will use this
 * identifier to refer to this match in the future
 * @param gameKey identifier of the game
 * @param locationKeys list of location identifiers, if the first one is not
 * available, the second one is tried and so on.
 * @param templateKey identifier of the template for this game to use
 * @param config configuration of the template
 */
func (client GameyeClient) CommandStartMatch(
	matchKey string,
	gameKey string,
	locationKeys []string,
	templateKey string,
	config map[string]interface{},
) (err error) {
	err = client.command("start-match", StartMatchPayload{
		matchKey,
		gameKey,
		locationKeys,
		templateKey,
		config,
	})
	return
}

type StopMatchPayload struct {
	MatchKey string
}

/**
 * Stop a match
 * @param matchKey Identifer of the match
 */
func (client GameyeClient) CommandStopMatch(
	matchKey string,
) (err error) {
	err = client.command("stop-match", StopMatchPayload{
		matchKey,
	})
	return
}

/**
 * Fetch the match state
 */
func (client GameyeClient) QueryMatch() (
	err error,
	state *models.MatchQueryState,
) {
	err = client.query(
		"match",
		nil,
		state,
	)
	return
}

/**
 * Subscribe to the match state
 */
func (client GameyeClient) SubscribeMatch(
	cancelChannel <-chan struct{},
) (
	err error,
	stateChannel chan<- *models.MatchQueryState,
) {
	var state *models.MatchQueryState
	var anyStateChannel <-chan interface{}

	anyStateChannel, err = client.subscribe(
		"match",
		nil,
		state,
		cancelChannel,
	)

	go func() {
		for anyState := range anyStateChannel {
			stateChannel <- anyState.(*models.MatchQueryState)
		}
	}()

	return
}
