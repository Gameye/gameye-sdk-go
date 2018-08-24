package models

var matchStateMock = &MatchQueryState{
	Match: MatchQueryMatchIndex{
		"test-match-123": &MatchQueryMatchItem{
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
		"test-match-456": &MatchQueryMatchItem{
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

/*
matchStateJSONMock is json mock data, useful for testing
*/
var matchStateJSONMock = `{
    "match": {
        "test-match-123": {
            "created": 1518191338368,
            "gameKey": "test",
            "host": "127.0.0.1",
            "locationKey": "local",
            "matchKey": "test-match-123",
            "port": {
                "game": 57015,
                "tv": 57025
            }
        },
        "test-match-456": {
            "created": 1518191339368,
            "gameKey": "testing",
            "host": "127.0.0.1",
            "locationKey": "local",
            "matchKey": "test-match-456",
            "port": {
                "game": 67015,
                "tv": 67025
            }
        }
    }
}`
