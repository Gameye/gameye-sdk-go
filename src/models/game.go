package models

type GameQueryState struct {
	Game     GameQueryGameIndex
	Location GameQueryLocationIndex
}

type GameQueryGameIndex = map[string]GameQueryGameItem

type GameQueryLocationIndex = map[string]GameQueryLocationItem

type GameQueryGameItem struct {
	GameKey  string
	Location map[string]bool
}
type GameQueryLocationItem struct {
	LocationKey string
}
