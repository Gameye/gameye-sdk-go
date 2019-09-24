package session

import (
	messages "github.com/Gameye/sdk-messages-go/pkg/event"
)

type OnData func(State)

var state State
var subscriptions map[string]OnData

func SubscribeState(name string, callback OnData) {
	if subscriptions == nil {
		subscriptions = make(map[string]OnData)
	}

	subscriptions[name] = callback
}

func UnsubscribeState(name string) {
	delete(subscriptions, name)
}

func Dispatch(action *messages.UnionEvent) {
	if action.Type != "" {
		state = reduce(&state, action)
		for _, callback := range subscriptions {
			callback(state)
		}
	}
}
