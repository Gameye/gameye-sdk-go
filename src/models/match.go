package models

type MatchQueryState struct {
	Match MatchQueryMatchIndex
}

type MatchQueryMatchIndex = map[string]*MatchQueryMatchItem

type MatchQueryMatchItem struct {
	MatchKey    string
	GameKey     string
	LocationKey string
	Host        string
	Created     int
	Port        map[string]int
}
