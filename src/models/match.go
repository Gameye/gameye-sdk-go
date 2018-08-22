package models

type MatchQueryState struct {
	Match MatchQueryMatchIndex `mapstructure:"match"`
}

type MatchQueryMatchIndex = map[string]MatchQueryMatchItem

type MatchQueryMatchItem struct {
	MatchKey    string         `mapstructure:"matchKey"`
	GameKey     string         `mapstructure:"gameKey"`
	LocationKey string         `mapstructure:"locationKey"`
	Host        string         `mapstructure:"host"`
	Created     int            `mapstructure:"created"`
	Port        map[string]int `mapstructure:"port"`
}
