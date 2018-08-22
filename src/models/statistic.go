package models

type StatisticQueryState struct {
	Statistic AnyStatisticState `mapstructure:"statistic"`
}

type AnyStatisticState = struct {
	Start          int                    `mapstructure:"start"`
	Stop           int                    `mapstructure:"stop"`
	StartedRounds  int                    `mapstructure:"startedRounds"`
	FinishedRounds int                    `mapstructure:"finishedRounds"`
	Player         map[string]PlayerModel `mapstructure:"player"`
	Team           map[string]TeamModel   `mapstructure:"team"`
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

type TeamModel struct {
	TeamKey   string          `mapstructure:"teamKey"`
	Name      string          `mapstructure:"name"`
	Statistic map[string]int  `mapstructure:"statistic"`
	Player    map[string]bool `mapstructure:"player"`
}
