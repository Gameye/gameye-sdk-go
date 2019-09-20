package session

import (
	"log"

	"github.com/mitchellh/mapstructure"

	messages "github.com/Gameye/sdk-messages-go/pkg/event"
)

func copyPorts(in *map[string]int64) (out map[string]int64) {
	for k, v := range *in {
		out[k] = v
	}
	return
}

func Reduce(state State, action *messages.UnionEvent) State {

	var sessions map[string]Session
	for k, v := range state.Sessions {
		sessions[k] = v
	}

	switch action.Type {
	case "session-initialized":
		var sessionInit messages.SessionInitializedEventPayload
		mapstructure.Decode(action.Payload, sessionInit)

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
		var sessionStarted messages.SessionStartedEventPayload
		mapstructure.Decode(action.Payload, sessionStarted)

		sessions[sessionStarted.Session.Id] = Session{
			Id:       sessionStarted.Session.Id,
			Image:    sessionStarted.Session.Image,
			Location: sessionStarted.Session.Location,
			Host:     sessionStarted.Session.Host,
			Created:  sessionStarted.Session.Created,
			Port:     copyPorts(&sessionStarted.Session.Port),
		}
	case "session-stopped":
		var sessionStopped messages.SessionStoppedEventPayload
		mapstructure.Decode(action.Payload, sessionStopped)
		log.Printf("%v\n", sessionStopped)
		delete(sessions, sessionStopped.Session.Id)
	}

	return StateWithSessions(sessions)
}
