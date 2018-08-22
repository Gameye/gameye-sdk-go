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
	cancelChannel <-chan struct{},
	matchKey string,
) (
	err error,
	stateChannel chan<- *models.StatisticQueryState,
) {
	var state *models.StatisticQueryState
	var anyStateChannel <-chan interface{}

	anyStateChannel, err = client.subscribe(
		"statistic",
		map[string]string{
			"matchKey": matchKey,
		},
		state,
		cancelChannel,
	)

	go func() {
		for anyState := range anyStateChannel {
			stateChannel <- anyState.(*models.StatisticQueryState)
		}
	}()

	return
}
