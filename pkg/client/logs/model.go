package logs

type LogLine struct {
	LineKey int    `json:"lineKey"`
	Payload string `json:"payload"`
}
