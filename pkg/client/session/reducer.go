package session

import (
	"encoding/json"
	"log"

	messages "github.com/Gameye/sdk-messages-go/pkg/event"
)

func copyPorts(in *map[string]int64) (out map[string]int64) {
	out = make(map[string]int64)
	for k, v := range *in {
		out[k] = v
	}
	return out
}

func reduce(state *State, action *messages.UnionEvent) State {

	sessions := make(map[string]Session)
	for k, v := range state.Sessions {
		sessions[k] = v
	}

	var err error

	switch action.Type {
	case "session-initialized":
		sessionInit := new(messages.SessionInitializedEventPayload)
		err = json.Unmarshal(*action.Payload, sessionInit)

		if err == nil {
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
		} else {
			log.Printf("session.reduce.session-initialized could not unmarshal; %v", err)
		}

	case "session-started":
		sessionStarted := new(messages.SessionStartedEventPayload)
		err = json.Unmarshal(*action.Payload, sessionStarted)

		if err == nil {
			sessions[sessionStarted.Session.Id] = Session{
				Id:       sessionStarted.Session.Id,
				Image:    sessionStarted.Session.Image,
				Location: sessionStarted.Session.Location,
				Host:     sessionStarted.Session.Host,
				Created:  sessionStarted.Session.Created,
				Port:     copyPorts(&sessionStarted.Session.Port),
			}
		} else {
			log.Printf("session.reduce.session-started could not unmarshal; %v", err)
		}

	case "session-stopped":
		sessionStopped := new(messages.SessionStoppedEventPayload)
		err = json.Unmarshal(*action.Payload, sessionStopped)

		if err == nil {
			delete(sessions, sessionStopped.Session.Id)
		} else {
			log.Printf("session.reduce.session-stopped could not unmarshal; %v", err)
		}
	}

	return StateWithSessions(sessions)
}
