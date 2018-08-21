package models

type StatisticQueryState struct {
	Statistic AnyStatisticState
}

type AnyStatisticState = struct {
	Start          int
	Stop           int
	StartedRounds  int
	FinishedRounds int
	Player         map[string]PlayerModel
	Team           map[string]int
}

type StartStopState struct {
	Start int
	Stop  int
}

type RoundState struct {
	StartedRounds  int
	FinishedRounds int
}

type PlayerContainerState struct {
	Player map[string]PlayerModel
}

type PlayerModel struct {
	PlayerKey string
	Uid       string
	Connected bool
	Name      string
	Statistic map[string]int
}

type TeamContainerState struct {
	Team map[string]TeamModel
}

type TeamModel struct {
	TeamKey   string
	Name      string
	Statistic map[string]int
	Player    map[string]bool
}
