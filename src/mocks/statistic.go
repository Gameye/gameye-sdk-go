package mocks

import "github.com/Gameye/gameye-sdk-go/src/models"

var StatisticStateMock = models.StatisticQueryState{
	Statistic: models.AnyStatisticState{
		Start: 1519833365000,
		Stop:  1519834524000,
		Player: map[string]models.PlayerModel{
			"3": models.PlayerModel{
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
			"4": models.PlayerModel{
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
		Team: map[string]models.TeamModel{
			"1": models.TeamModel{
				TeamKey: "1",
				Name:    "TeamA",
				Statistic: map[string]int{
					"score": 17,
				},
				Player: map[string]bool{
					"3": true,
				},
			},
			"2": models.TeamModel{
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
