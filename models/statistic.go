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

type StatisticQueryState struct {
	Statistic AnyStatisticState `mapstructure:"statistic"`
}

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

var StatisticStateMock = StatisticQueryState{
	Statistic: AnyStatisticState{
		Start: 1519833365000,
		Stop:  1519834524000,
		Player: map[string]*PlayerModel{
			"3": &PlayerModel{
				PlayerKey: "3",
				Connected: false,
				UID:       "STEAM_1:1:218909830",
				Name:      "denise",
				Statistic: map[string]int{
					"assist": 0,
					"death":  19,
					"kill":   17,
				},
			},
			"4": &PlayerModel{
				PlayerKey: "4",
				Connected: false,
				UID:       "STEAM_1:1:24748064",
				Name:      "Smashmint",
				Statistic: map[string]int{
					"assist": 0,
					"death":  17,
					"kill":   19,
				},
			},
		},
		StartedRounds:  36,
		FinishedRounds: 36,
		Team: map[string]*TeamModel{
			"1": &TeamModel{
				TeamKey: "1",
				Name:    "TeamA",
				Statistic: map[string]int{
					"score": 17,
				},
				Player: map[string]bool{
					"3": true,
				},
			},
			"2": &TeamModel{
				TeamKey: "2",
				Name:    "TeamB",
				Statistic: map[string]int{
					"score": 19,
				},
				Player: map[string]bool{
					"4": true,
				},
			},
		},
	},
}

var StatisticStateJSONMock = `{
    "statistic": {
        "start": 1519833365000,
        "stop": 1519834524000,
        "player": {
            "3": {
                "playerKey": "3",
                "connected": false,
                "uid": "STEAM_1:1:218909830",
                "name": "denise",
                "statistic": {
                    "assist": 0,
                    "death": 19,
                    "kill": 17
                }
            },
            "4": {
                "playerKey": "4",
                "connected": false,
                "uid": "STEAM_1:1:24748064",
                "name": "Smashmint",
                "statistic": {
                    "assist": 0,
                    "death": 17,
                    "kill": 19
                }
            }
        },
        "startedRounds": 36,
        "finishedRounds": 36,
        "team": {
            "1": {
                "teamKey": "1",
                "name": "TeamA",
                "statistic": {
                    "score": 17
                },
                "player": {
                    "3": true
                }
            },
            "2": {
                "teamKey": "2",
                "name": "TeamB",
                "statistic": {
                    "score": 19
                },
                "player": {
                    "4": true
                }
            }
        }
    }
}`
