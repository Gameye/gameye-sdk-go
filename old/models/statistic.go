package models

import "github.com/mitchellh/mapstructure"

/*
CreateStatisticQueryState will create a new StatisticQueryState from
a string / interface map
*/
func CreateStatisticQueryState(anyState *map[string]interface{}) (
	state *StatisticQueryState,
	err error,
) {
	state = &StatisticQueryState{}
	mapstructure.Decode(anyState, state)
	return
}

/*
StatisticQueryState is data coming from the api
*/
type StatisticQueryState struct {
	Statistic AnyStatisticState `mapstructure:"statistic"`
}

/*
AnyStatisticState is data coming from the api
*/
type AnyStatisticState = struct {
	Start          int                     `mapstructure:"start"`
	Stop           int                     `mapstructure:"stop"`
	StartedRounds  int                     `mapstructure:"startedRounds"`
	FinishedRounds int                     `mapstructure:"finishedRounds"`
	Player         map[string]*PlayerModel `mapstructure:"player"`
	Team           map[string]*TeamModel   `mapstructure:"team"`
}

// type StartStopState struct {
// 	Start int `mapstructure:"start"`
// 	Stop  int `mapstructure:"stop"`
// }

// type RoundState struct {
// 	StartedRounds  int `mapstructure:"startedRounds"`
// 	FinishedRounds int `mapstructure:"finishedRounds"`
// }

// type PlayerContainerState struct {
// 	Player map[string]PlayerModel `mapstructure:"player"`
// }

/*
PlayerModel is data coming from the api
*/
type PlayerModel struct {
	PlayerKey string         `mapstructure:"playerKey"`
	UID       string         `mapstructure:"uid"`
	Connected bool           `mapstructure:"connected"`
	Name      string         `mapstructure:"name"`
	Statistic map[string]int `mapstructure:"statistic"`
}

// type TeamContainerState struct {
// 	Team map[string]int `mapstructure:"team"`
// }

/*
TeamModel is data coming from the api
*/
type TeamModel struct {
	TeamKey   string          `mapstructure:"teamKey"`
	Name      string          `mapstructure:"name"`
	Statistic map[string]int  `mapstructure:"statistic"`
	Player    map[string]bool `mapstructure:"player"`
}
