package event

// SessionStartedEvent is received when a session is started
type SessionStartedEvent struct {
	Type    string                     `json:"type"`
	Payload SessionStartedEventPayload `json:"payload"`
}

// SessionStartedEventPayload is the payload for the SessionStartedEvent
type SessionStartedEventPayload struct {
	Session SessionStarted `json:"session"`
}

// SessionStarted contains the information about a started session
type SessionStarted struct {
	ID       string           `json:"id"`
	Image    string           `json:"image"`
	Location string           `json:"location"`
	Host     string           `json:"host"`
	Created  int64            `json:"created"`
	Port     map[string]int64 `json:"port"`
}
