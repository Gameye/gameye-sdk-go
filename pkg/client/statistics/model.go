package statistics

// Player represents a player in the Session, identified by PlayerKey
type Player struct {
	PlayerKey string             `json:"playerKey"`
	UID       string             `json:"uid"`
	Connected bool               `json:"connected"`
	Name      string             `json:"name"`
	Statistic map[string]float64 `json:"statistic"`
}

// Team represents a team in the Session, identified by TeamKey
type Team struct {
	TeamKey   string             `json:"teamKey"`
	Name      string             `json:"name"`
	Statistic map[string]float64 `json:"statistic"`
	Player    map[string]bool    `json:"player"`
}
