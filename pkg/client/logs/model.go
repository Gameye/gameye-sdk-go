package logs

// LogLine represents a line number and log message
type LogLine struct {
	LineKey int    `json:"lineKey"`
	Payload string `json:"payload"`
}
