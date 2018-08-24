package clients

import (
	"github.com/Gameye/gameye-sdk-go/models"
)

/*
QueryTemplate will fetch template state
@param gameKey identifier of the game
*/
func (client GameyeClient) QueryTemplate(
	gameKey string,
) (
	state *models.TemplateQueryState,
	err error,
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

	state, err = models.CreateTemplateQueryState(&anyState)
	if err != nil {
		return
	}

	return
}

/*
SubscribeTemplate will subscribe to template state
@param gameKey identifier of the game
*/
func (client GameyeClient) SubscribeTemplate(
	gameKey string,
) (
	subscription *TemplateQuerySubscription,
	err error,
) {
	var qs *querySubscription
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

/*
TemplateQuerySubscription is a subscription to the template state
*/
type TemplateQuerySubscription struct {
	qs *querySubscription
}

/*
Cancel will end and cleanup the subscription
*/
func (s *TemplateQuerySubscription) Cancel() {
	s.qs.Cancel()
}

/*
NextState will return the next state
*/
func (s *TemplateQuerySubscription) NextState() (
	state *models.TemplateQueryState,
	err error,
) {
	var anyState map[string]interface{}
	anyState, err = s.qs.NextState()
	if err != nil {
		return
	}
	state, err = models.CreateTemplateQueryState(&anyState)
	if err != nil {
		return
	}

	return
}
