package event

// SessionStoppedEvent is received when a session is ended
type SessionStoppedEvent struct {
	Type    string                     `json:"type"`
	Payload SessionStoppedEventPayload `json:"payload"`
}

// SessionStoppedEventPayload is the payload for the SessionStoppedEvent
type SessionStoppedEventPayload struct {
	Session SessionStopped `json:"session"`
}

// SessionStopped contains the identifier of a session
type SessionStopped struct {
	ID string `json:"id"`
}
