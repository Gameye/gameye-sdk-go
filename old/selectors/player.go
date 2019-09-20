package selectors

import "github.com/Gameye/gameye-sdk-go/models"

/*
PlayerItem holds information about a player
*/
type PlayerItem = models.PlayerModel

/*
SelectPlayerList lists all players in the match.
@param statisticState statistic state
*/
func SelectPlayerList(
	statisticState *models.StatisticQueryState,
) (
	playerList []*PlayerItem,
) {
	playerList = make([]*PlayerItem, 0)

	playerIndex := statisticState.Statistic.Player
	if playerIndex == nil {
		return
	}

	for _, playerItem := range playerIndex {
		if playerItem == nil {
			continue
		}

		playerList = append(playerList, playerItem)
	}

	return
}

/*
SelectPlayerListForTeam gets a list if all players in a team.
@param statisticState statistic state
@param teamKey identifier of the team
*/
func SelectPlayerListForTeam(
	statisticState *models.StatisticQueryState,
	teamKey string,
) (
	playerList []*PlayerItem,
) {
	playerList = make([]*PlayerItem, 0)

	teamIndex := statisticState.Statistic.Team
	if teamIndex == nil {
		return
	}

	teamItem := teamIndex[teamKey]
	if teamItem == nil {
		return
	}

	playerIndex := statisticState.Statistic.Player
	if playerIndex == nil {
		return
	}

	for playerKey, playerEnabled := range teamItem.Player {
		if !playerEnabled {
			continue
		}

		playerItem := playerIndex[playerKey]
		if playerItem == nil {
			continue
		}

		playerList = append(playerList, playerItem)
	}

	return
}

/*
SelectPlayerItem gets a single player in the match.
@param statisticState statistic state
@param playerKey identifier of the player to get the details for
*/
func SelectPlayerItem(
	statisticState *models.StatisticQueryState,
	playerKey string,
) (
	playerItem *PlayerItem,
) {
	playerIndex := statisticState.Statistic.Player
	if playerIndex == nil {
		return
	}

	playerItem = playerIndex[playerKey]
	return
}
