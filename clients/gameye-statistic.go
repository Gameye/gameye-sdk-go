package clients

import (
	"github.com/Gameye/gameye-sdk-go/models"
)

/*
QueryStatistic will fetch statistic state
@param matchKey identifier of the match
*/
func (client GameyeClient) QueryStatistic(
	matchKey string,
) (
	state *models.StatisticQueryState,
	err error,
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

/*
SubscribeStatistic will subscribe to statistic state
@param matchKey identifier of the match
*/
func (client GameyeClient) SubscribeStatistic(
	matchKey string,
) (
	subscription *StatisticQuerySubscription,
	err error,
) {
	var qs *querySubscription
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

/*
StatisticQuerySubscription is a subscription to the statistic state
*/
type StatisticQuerySubscription struct {
	qs *querySubscription
}

/*
Cancel will end and cleanup the subscription
*/
func (s *StatisticQuerySubscription) Cancel() {
	s.qs.Cancel()
}

/*
NextState will return the next state
*/
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
