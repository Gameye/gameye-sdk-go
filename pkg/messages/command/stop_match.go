package command

// StopMatchCommand is used to stop a match
type StopMatchCommand struct {
	Type    string                  `json:"type"`
	Payload StopMatchCommandPayload `json:"payload"`
}

// StopMatchCommandPayload is the payload for the StopMatchCommand
type StopMatchCommandPayload struct {
	AccountKey string `json:"accountKey"`
	MatchKey   string `json:"matchKey"`
}
