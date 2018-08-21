package selectors

import "github.com/Gameye/gameye-sdk-go/src/models"

type PlayerItem = models.PlayerModel

/**
 * List all players in the match.
 * @param statisticState statistic state
 */
func SelectPlayerList(
	statisticState *models.StatisticQueryState,
) (
	playerList []*PlayerItem,
) {
	// const playerIndex = statisticState.statistic.player;
	// if (!playerIndex) return [];

	// return Object.values(playerIndex);
	return
}

/**
 * Get a list if all players in a team.
 * @param statisticState statistic state
 * @param teamKey identifier of the team
 */
func SelectPlayerListForTeam(
	statisticState *models.StatisticQueryState,
	teamKey string,
) (
	playerList []*PlayerItem,
) {
	// const teamIndex = statisticState.statistic.team;
	// if (!teamIndex) return [];

	// if (!teamIndex[teamKey]) return [];

	// const playerIndex = statisticState.statistic.player;
	// if (!playerIndex) return [];

	// return Object.entries(teamIndex[teamKey].player).
	//     filter(([, playerEnabled]) => playerEnabled).
	//     map(([playerKey]) => playerIndex[playerKey]);
	return
}

/**
 * Get a single player in the match.
 * @param statisticState statistic state
 * @param playerKey identifier of the player to get the details for
 */
func SelectPlayerItem(
	statisticState *models.StatisticQueryState,
	playerKey string,
) (
	playerItem *PlayerItem,
) {
	// const playerIndex = statisticState.statistic.player;
	// if (!playerIndex) return null;

	// const playerItem = playerIndex[playerKey];
	// if (!playerItem) return null;

	// return playerItem;
	return
}
