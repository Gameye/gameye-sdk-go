package models

import "github.com/mitchellh/mapstructure"

/*
CreateMatchQueryState will create a new MatchQueryState from
a string / interface map
*/
func CreateMatchQueryState(anyState *map[string]interface{}) (
	state *MatchQueryState,
	err error,
) {
	state = &MatchQueryState{}
	mapstructure.Decode(anyState, state)
	return
}

/*
MatchQueryState is data coming from the api
*/
type MatchQueryState struct {
	Match MatchQueryMatchIndex `mapstructure:"match"`
}

/*
MatchQueryMatchIndex is data coming from the api
*/
type MatchQueryMatchIndex = map[string]*MatchQueryMatchItem

/*
MatchQueryMatchItem is data coming from the api
*/
type MatchQueryMatchItem struct {
	MatchKey    string         `mapstructure:"matchKey"`
	GameKey     string         `mapstructure:"gameKey"`
	LocationKey string         `mapstructure:"locationKey"`
	Host        string         `mapstructure:"host"`
	Created     int            `mapstructure:"created"`
	Port        map[string]int `mapstructure:"port"`
}

/*
MatchStateMock is mock data, useful for testing
*/
var MatchStateMock = MatchQueryState{
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
MatchStateJSONMock is json mock data, useful for testing
*/
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
