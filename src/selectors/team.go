package selectors

import "github.com/Gameye/gameye-sdk-go/src/models"

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
	// const teamIndex = statisticState.statistic.team;
	// if (!teamIndex) return [];

	// return Object.values(teamIndex);
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
	// const teamIndex = statisticState.statistic.team;
	// if (!teamIndex) return null;

	// const teamItem = teamIndex[teamKey];
	// if (!teamItem) return null;

	// return teamItem;
	return
}
