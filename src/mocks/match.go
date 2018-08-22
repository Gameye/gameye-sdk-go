package mocks

import "github.com/Gameye/gameye-sdk-go/src/models"

var MatchStateMock = models.MatchQueryState{
	Match: models.MatchQueryMatchIndex{
		"test-match-123": models.MatchQueryMatchItem{
			Created:     1518191338368,
			GameKey:     "test",
			Host:        "127.0.0.1",
			LocationKey: "local",
			MatchKey:    "test-match-123",
			Port: map[string]int{
				"game": 57015,
				"tv":   57025,
			},
		},
		"test-match-456": models.MatchQueryMatchItem{
			Created:     1518191339368,
			GameKey:     "testing",
			Host:        "127.0.0.1",
			LocationKey: "local",
			MatchKey:    "test-match-456",
			Port: map[string]int{
				"game": 67015,
				"tv":   67025,
			},
		},
	},
}

var MatchStateJSONMock = `{
    "match": {
        "test-match-123": {
            "created": 1518191338368,
            "gameKey": "test",
            "host": "127.0.0.1",
            "locationKey": "local",
            "matchKey": "test-match-123",
            "port": {
                "game": 57015,
                "tv": 57025,
            },
        },
        "test-match-456": {
            "created": 1518191339368,
            "gameKey": "testing",
            "host": "127.0.0.1",
            "locationKey": "local",
            "matchKey": "test-match-456",
            "port": {
                "game": 67015,
                "tv": 67025,
            },
        },
    },
}`
