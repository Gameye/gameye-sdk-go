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
