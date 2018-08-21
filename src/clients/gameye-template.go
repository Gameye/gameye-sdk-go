package clients

import (
	"github.com/Gameye/gameye-sdk-go/src/models"
)

type TemplateQueryArg struct {
	GameKey string
}

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
	client.query("template", TemplateQueryArg{
		gameKey,
	})
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
	state *models.TemplateQueryState,
) {
	client.subscribe("template", TemplateQueryArg{
		gameKey,
	})
	return
}
