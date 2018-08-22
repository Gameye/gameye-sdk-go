package selectors

import "github.com/Gameye/gameye-sdk-go/models"

type LocationItem = models.GameQueryLocationItem

/**
 * Selects all locations for a given game.
 * @param gameState game state
 * @param gameKey identifier of the game
 */
func SelectLocationListForGame(
	gameState *models.GameQueryState,
	gameKey string,
) (
	locationList []*LocationItem,
) {
	locationList = make([]*LocationItem, 0)

	gameItem := gameState.Game[gameKey]
	if gameItem == nil {
		return
	}

	for locationKey, hasLocation := range gameItem.Location {
		if !hasLocation {
			continue
		}
		locationItem := gameState.Location[locationKey]
		locationList = append(locationList, locationItem)
	}

	return
}
