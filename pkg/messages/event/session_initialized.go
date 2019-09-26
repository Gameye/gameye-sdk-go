package event

// SessionInitializedEvent is received when a session event stream is opened
type SessionInitializedEvent struct {
	Type    string                         `json:"type"`
	Payload SessionInitializedEventPayload `json:"payload"`
}

// SessionInitializedEventPayload is the payload for the SessionInitializedEvent
type SessionInitializedEventPayload struct {
	Sessions []SessionInitialized `json:"sessions"`
}

// SessionInitialized contains the information about an active sessions
type SessionInitialized struct {
	ID       string           `json:"id"`
	Image    string           `json:"image"`
	Location string           `json:"location"`
	Host     string           `json:"host"`
	Created  int64            `json:"created"`
	Port     map[string]int64 `json:"port"`
}
