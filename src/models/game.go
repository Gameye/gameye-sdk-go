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
