package clients

import (
	"github.com/Gameye/gameye-sdk-go/models"
	"github.com/mitchellh/mapstructure"
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
	var anyState map[string]interface{}

	anyState, err = client.query(
		"template",
		map[string]string{
			"gameKey": gameKey,
		},
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
 * Subscribe to template state
 * @param gameKey identifier of the game
 */
func (client GameyeClient) SubscribeTemplate(
	gameKey string,
) (
	err error,
	subscription *TemplateQuerySubscription,
) {
	var qs QuerySubscription
	qs, err = client.subscribe(
		"template",
		map[string]string{
			"gameKey": gameKey,
		},
	)
	if err != nil {
		return
	}

	subscription = &TemplateQuerySubscription{
		qs,
	}

	return
}

type TemplateQuerySubscription struct {
	qs QuerySubscription
}

func (s *TemplateQuerySubscription) Cancel() {
	s.qs.Cancel()
}

func (s *TemplateQuerySubscription) NextState() (
	state *models.TemplateQueryState,
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
