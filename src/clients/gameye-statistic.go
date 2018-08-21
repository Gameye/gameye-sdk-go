package clients

import (
	"github.com/Gameye/gameye-sdk-go/src/models"
)

type StatisticQueryArg struct {
	MatchKey string
}

/**
 * Fetch statistic state
 * @param matchKey identifier of the match
 */
func (client GameyeClient) QueryStatistic(
	matchKey string,
) (
	err error,
	state *models.StatisticQueryState,
) {
	err = client.query(
		"statistic",
		map[string]string{
			"matchKey": matchKey,
		},
		state,
	)
	return
}

/**
 * Subscribe to statistic state
 * @param matchKey identifier of the match
 */
func (client GameyeClient) SubscribeStatistic(
	matchKey string,
) (
	err error,
	state *models.StatisticQueryState,
) {
	err = client.subscribe(
		"statistic",
		map[string]string{
			"matchKey": matchKey,
		},
		state,
	)
	return
}
