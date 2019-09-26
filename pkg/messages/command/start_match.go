package command

// StartMatchCommand is used to start a match
type StartMatchCommand struct {
	Type    string                   `json:"type"`
	Payload StartMatchCommandPayload `json:"payload"`
}

// StartMatchCommandPayload is the payload for the StartMatchCommand
type StartMatchCommandPayload struct {
	AccountKey     string                 `json:"accountKey"`
	MatchKey       string                 `json:"matchKey"`
	GameKey        string                 `json:"gameKey"`
	LocationKeys   []string               `json:"locationKeys"`
	TemplateKey    string                 `json:"templateKey"`
	Config         map[string]interface{} `json:"config"`
	EndCallbackURL string                 `json:"endCallbackUrl"`
}
