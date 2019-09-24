package session

// State represents all the session data
type State struct {
	Sessions map[string]Session
}

// StateWithSessions constructs a State object with a set of sessions
func StateWithSessions(sessions map[string]Session) (state State) {
	state = State{
		sessions,
	}
	return
}
