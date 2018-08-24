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
