package statistics


type Player struct
{
	PlayerKey string `json:"playerKey"`
	Uid string `json:"uid"`
	Connected bool `json:"connected"`
	Name string `json:"name"`
	Statistic map[string]float64 `json:"statistic"`
}

type Team struct
{
	TeamKey string `json:"teamKey"`
	Name string `json:"name"`
	Statistic map[string]float64 `json:"statistic"`
	Player map[string]bool `json:"player"`
}