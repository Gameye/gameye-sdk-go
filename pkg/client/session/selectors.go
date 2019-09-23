package session

func SelectSessionListForGame(state State, gameKey string) (sessions []Session) {
	for _, session := range state.Sessions {
		if session.Image == gameKey {
			sessions = append(sessions, session)
		}
	}
	return sessions
}

func SelectSessionList(state State) (sessions []Session) {
	for _, session := range state.Sessions {
		sessions = append(sessions, session)
	}
	return sessions
}

func SelectSession(state State, sessionId string) (session Session) {
	return state.Sessions[sessionId]
}
