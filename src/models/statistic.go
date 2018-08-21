package models

type StatisticQueryState struct {
	statistic AnyStatisticState
}

type AnyStatisticState = interface{}

type StartStopState struct {
	Start int
	Stop  int
}

type RoundState struct {
	StartedRounds  int
	FinishedRounds int
}

type PlayerContainerState struct {
	Player map[string]*PlayerModel
}

type PlayerModel struct {
	PlayerKey string
	Uid       string
	Connected bool
	Name      string
	Statistic map[string]int
}

type TeamContainerState struct {
	Team map[string]int
}

type TeamModel struct {
	TeamKey   string
	Name      string
	Statistic map[string]int
	Player    map[string]bool
}
