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
	gameKey string,
) (
	err error,
	state *models.TemplateQueryState,
) {
	err = client.subscribe(
		"template",
		map[string]string{
			"gameKey": gameKey,
		},
		state,
	)
	return
}
