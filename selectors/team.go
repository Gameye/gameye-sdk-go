package selectors

import "github.com/Gameye/gameye-sdk-go/models"

type TeamItem = models.TeamModel

/**
 * Get a list of all teams in the statistic-state.
 * @param statisticState statistic state
 */
func SelectTeamList(
	statisticState *models.StatisticQueryState,
) (
	teamList []*TeamItem,
) {
	teamList = make([]*TeamItem, 0)

	teamIndex := statisticState.Statistic.Team
	if teamIndex == nil {
		return
	}

	for _, teamItem := range teamIndex {
		if teamItem == nil {
			continue
		}

		teamList = append(teamList, teamItem)
	}

	return
}

/**
 * Get a single team from the statistic-state.
 * @param statisticState statistic state
 * @param teamKey identifier of the team
 */
func SelectTeamItem(
	statisticState *models.StatisticQueryState,
	teamKey string,
) (
	teamItem *TeamItem,
) {
	teamIndex := statisticState.Statistic.Team
	if teamIndex == nil {
		return
	}

	teamItem = teamIndex[teamKey]
	return
}
