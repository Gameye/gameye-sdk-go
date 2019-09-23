package session

import (
	"encoding/json"

	messages "github.com/Gameye/sdk-messages-go/pkg/event"
)

func copyPorts(in *map[string]int64) (out map[string]int64) {
	out = make(map[string]int64)
	for k, v := range *in {
		out[k] = v
	}
	return out
}

func Reduce(state *State, action *messages.UnionEvent) State {

	sessions := make(map[string]Session)

	for k, v := range state.Sessions {
		sessions[k] = v
	}

	switch action.Type {
	case "session-initialized":
		sessionInit := new(messages.SessionInitializedEventPayload)
		json.Unmarshal(*action.Payload, sessionInit)

		for _, session := range sessionInit.Sessions {
			sessions[session.Id] = Session{
				Id:       session.Id,
				Image:    session.Image,
				Location: session.Location,
				Host:     session.Host,
				Created:  session.Created,
				Port:     copyPorts(&session.Port),
			}
		}
	case "session-started":
		sessionStarted := new(messages.SessionStartedEventPayload)
		json.Unmarshal(*action.Payload, sessionStarted)

		sessions[sessionStarted.Session.Id] = Session{
			Id:       sessionStarted.Session.Id,
			Image:    sessionStarted.Session.Image,
			Location: sessionStarted.Session.Location,
			Host:     sessionStarted.Session.Host,
			Created:  sessionStarted.Session.Created,
			Port:     copyPorts(&sessionStarted.Session.Port),
		}
	case "session-stopped":
		sessionStopped := new(messages.SessionStoppedEventPayload)
		json.Unmarshal(*action.Payload, sessionStopped)

		delete(sessions, sessionStopped.Session.Id)
	}

	return StateWithSessions(sessions)
}
