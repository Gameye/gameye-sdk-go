package models

import "github.com/mitchellh/mapstructure"

/*
CreateGameQueryState will create a new GameQueryState from
a string / interface map
*/
func CreateGameQueryState(anyState *map[string]interface{}) (
	state *GameQueryState,
	err error,
) {
	state = &GameQueryState{}
	mapstructure.Decode(anyState, state)
	return
}

/*
GameQueryState is data coming from the api
*/
type GameQueryState struct {
	Game     GameQueryGameIndex     `mapstructure:"game"`
	Location GameQueryLocationIndex `mapstructure:"location"`
}

/*
GameQueryGameIndex is data coming from the api
*/
type GameQueryGameIndex = map[string]*GameQueryGameItem

/*
GameQueryLocationIndex is data coming from the api
*/
type GameQueryLocationIndex = map[string]*GameQueryLocationItem

/*
GameQueryGameItem is data coming from the api
*/
type GameQueryGameItem struct {
	GameKey  string          `mapstructure:"gameKey"`
	Location map[string]bool `mapstructure:"location"`
}

/*
GameQueryLocationItem is data coming from the api
*/
type GameQueryLocationItem struct {
	LocationKey string `mapstructure:"locationKey"`
}

/*
GameStateMock is mock data, useful for testing
*/
var GameStateMock = GameQueryState{
	Game: GameQueryGameIndex{
		"csgo": &GameQueryGameItem{GameKey: "csgo", Location: map[string]bool{}},
		"tf2":  &GameQueryGameItem{GameKey: "tf2", Location: map[string]bool{}},
		"css":  &GameQueryGameItem{GameKey: "css", Location: map[string]bool{}},
		"l4d2": &GameQueryGameItem{GameKey: "l4d2", Location: map[string]bool{}},
		"kf2":  &GameQueryGameItem{GameKey: "kf2", Location: map[string]bool{}},
		"test": &GameQueryGameItem{GameKey: "test", Location: map[string]bool{"local": true}},
	},
	Location: GameQueryLocationIndex{
		"rotterdam":     &GameQueryLocationItem{LocationKey: "rotterdam"},
		"ireland":       &GameQueryLocationItem{LocationKey: "ireland"},
		"dubai":         &GameQueryLocationItem{LocationKey: "dubai"},
		"tokyo":         &GameQueryLocationItem{LocationKey: "tokyo"},
		"los_angeles":   &GameQueryLocationItem{LocationKey: "los_angeles"},
		"washington_dc": &GameQueryLocationItem{LocationKey: "washington_dc"},
		"local":         &GameQueryLocationItem{LocationKey: "local"},
	},
}

/*
GameStateJSONMock is json mock data, useful for testing
*/
var GameStateJSONMock = `{
    "game": {
        "csgo": { "gameKey": "csgo", "location": {} },
        "tf2": { "gameKey": "tf2", "location": {} },
        "css": { "gameKey": "css", "location": {} },
        "l4d2": { "gameKey": "l4d2", "location": {} },
        "kf2": { "gameKey": "kf2", "location": {} },
        "test": { "gameKey": "test", "location": { "local": true } }
    },
    "location": {
        "rotterdam": { "locationKey": "rotterdam" },
        "ireland": { "locationKey": "ireland" },
        "dubai": { "locationKey": "dubai" },
        "tokyo": { "locationKey": "tokyo" },
        "los_angeles": { "locationKey": "los_angeles" },
        "washington_dc": { "locationKey": "washington_dc" },
        "local": { "locationKey": "local" }
    }
}`
