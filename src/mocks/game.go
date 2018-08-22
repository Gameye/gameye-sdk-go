package mocks

import "github.com/Gameye/gameye-sdk-go/src/models"

var GameStateMock = models.GameQueryState{
	Game: models.GameQueryGameIndex{
		"csgo": models.GameQueryGameItem{GameKey: "csgo", Location: map[string]bool{}},
		"tf2":  models.GameQueryGameItem{GameKey: "tf2", Location: map[string]bool{}},
		"css":  models.GameQueryGameItem{GameKey: "css", Location: map[string]bool{}},
		"l4d2": models.GameQueryGameItem{GameKey: "l4d2", Location: map[string]bool{}},
		"kf2":  models.GameQueryGameItem{GameKey: "kf2", Location: map[string]bool{}},
		"test": models.GameQueryGameItem{GameKey: "test", Location: map[string]bool{"local": true}},
	},
	Location: models.GameQueryLocationIndex{
		"rotterdam":     models.GameQueryLocationItem{LocationKey: "rotterdam"},
		"ireland":       models.GameQueryLocationItem{LocationKey: "ireland"},
		"dubai":         models.GameQueryLocationItem{LocationKey: "dubai"},
		"tokyo":         models.GameQueryLocationItem{LocationKey: "tokyo"},
		"los_angeles":   models.GameQueryLocationItem{LocationKey: "los_angeles"},
		"washington_dc": models.GameQueryLocationItem{LocationKey: "washington_dc"},
		"local":         models.GameQueryLocationItem{LocationKey: "local"},
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
