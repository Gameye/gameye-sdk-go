package clients

import (
	"github.com/Gameye/gameye-sdk-go/models"
)

// type StatisticQueryArg struct {
// 	MatchKey string
// }

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
	var anyState map[string]interface{}

	anyState, err = client.query(
		"statistic",
		map[string]string{
			"matchKey": matchKey,
		},
	)
	if err != nil {
		return
	}

	state, err = models.CreateStatisticQueryState(&anyState)
	if err != nil {
		return
	}

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
	subscription *StatisticQuerySubscription,
) {
	var qs QuerySubscription
	qs, err = client.subscribe(
		"statistic",
		map[string]string{
			"matchKey": matchKey,
		},
	)

	if err != nil {
		return
	}

	subscription = &StatisticQuerySubscription{
		qs,
	}

	return
}

type StatisticQuerySubscription struct {
	qs QuerySubscription
}

func (s *StatisticQuerySubscription) Cancel() {
	s.qs.Cancel()
}

func (s *StatisticQuerySubscription) NextState() (
	state *models.StatisticQueryState,
	err error,
) {
	var anyState map[string]interface{}
	anyState, err = s.qs.NextState()
	if err != nil {
		return
	}
	state, err = models.CreateStatisticQueryState(&anyState)
	if err != nil {
		return
	}

	return
}
