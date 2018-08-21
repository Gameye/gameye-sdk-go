package selectors

import "github.com/Gameye/gameye-sdk-go/src/models"

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
	// const gameItem = gameState.game[gameKey];
	// if (!gameItem) return [];

	// return Object.entries(gameItem.location).
	//     filter(([, hasLocation]) => hasLocation).
	//     map(([locationKey]) => ({
	//         locationKey,
	//     } as LocationItem));

	return
}
