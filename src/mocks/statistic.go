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
				Uid:       "STEAM_1:1:218909830",
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
				Uid:       "STEAM_1:1:24748064",
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
				teamKey: "1",
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
