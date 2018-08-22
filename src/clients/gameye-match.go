package clients

import (
	"github.com/Gameye/gameye-sdk-go/src/models"
	"github.com/mitchellh/mapstructure"
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
	var anyState map[string]interface{}

	anyState, err = client.query(
		"match",
		nil,
	)
	if err != nil {
		return
	}

	err = mapstructure.Decode(anyState, state)
	if err != nil {
		return
	}

	return
}

/**
 * Subscribe to the match state
 */
func (client GameyeClient) SubscribeMatch() (
	err error,
	subscription *MatchQuerySubscription,
) {
	var qs QuerySubscription
	qs, err = client.subscribe(
		"match",
		nil,
	)
	if err != nil {
		return
	}

	subscription = &MatchQuerySubscription{
		qs,
	}

	return
}

type MatchQuerySubscription struct {
	qs QuerySubscription
}

func (s *MatchQuerySubscription) Cancel() {
	s.qs.Cancel()
}

func (s *MatchQuerySubscription) NextState() (
	state *models.MatchQueryState,
	err error,
) {
	var anyState map[string]interface{}
	anyState, err = s.qs.NextState()
	if err != nil {
		return
	}
	err = mapstructure.Decode(anyState, state)
	if err != nil {
		return
	}

	return
}
