package session

// Session represents a game session identified by ID
type Session struct {
	ID       string           `json:"id"`
	Image    string           `json:"image"`
	Location string           `json:"location"`
	Host     string           `json:"host"`
	Created  int64            `json:"created"`
	Port     map[string]int64 `json:"port"`
}
