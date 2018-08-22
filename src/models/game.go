package models

type GameQueryState struct {
	Game     GameQueryGameIndex     `mapstructure:"game"`
	Location GameQueryLocationIndex `mapstructure:"location"`
}

type GameQueryGameIndex = map[string]GameQueryGameItem

type GameQueryLocationIndex = map[string]GameQueryLocationItem

type GameQueryGameItem struct {
	GameKey  string          `mapstructure:"gameKey"`
	Location map[string]bool `mapstructure:"location"`
}
type GameQueryLocationItem struct {
	LocationKey string `mapstructure:"locationKey"`
}

var GameStateMock = GameQueryState{
	Game: GameQueryGameIndex{
		"csgo": GameQueryGameItem{GameKey: "csgo", Location: map[string]bool{}},
		"tf2":  GameQueryGameItem{GameKey: "tf2", Location: map[string]bool{}},
		"css":  GameQueryGameItem{GameKey: "css", Location: map[string]bool{}},
		"l4d2": GameQueryGameItem{GameKey: "l4d2", Location: map[string]bool{}},
		"kf2":  GameQueryGameItem{GameKey: "kf2", Location: map[string]bool{}},
		"test": GameQueryGameItem{GameKey: "test", Location: map[string]bool{"local": true}},
	},
	Location: GameQueryLocationIndex{
		"rotterdam":     GameQueryLocationItem{LocationKey: "rotterdam"},
		"ireland":       GameQueryLocationItem{LocationKey: "ireland"},
		"dubai":         GameQueryLocationItem{LocationKey: "dubai"},
		"tokyo":         GameQueryLocationItem{LocationKey: "tokyo"},
		"los_angeles":   GameQueryLocationItem{LocationKey: "los_angeles"},
		"washington_dc": GameQueryLocationItem{LocationKey: "washington_dc"},
		"local":         GameQueryLocationItem{LocationKey: "local"},
	},
}

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
