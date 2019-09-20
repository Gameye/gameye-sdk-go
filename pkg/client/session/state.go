package session

type State struct {
	Sessions map[string]Session
}

func StateWithSessions(sessions map[string]Session) (state State) {
	state = State{
		sessions,
	}
	return
}
