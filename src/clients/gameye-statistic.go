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
	client.query("statistic", StatisticQueryArg{
		matchKey,
	})
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
	client.subscribe("statistic", StatisticQueryArg{
		matchKey,
	})
	return
}
