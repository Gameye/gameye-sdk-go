package session

import (
	messages "../../messages/event"
)

// OnData is a callback type used for events
type OnData func(State)

var state State
var subscriptions map[string]OnData

// SubscribeState triggers callback events when the state changes
func SubscribeState(name string, callback OnData) {
	if subscriptions == nil {
		subscriptions = make(map[string]OnData)
	}

	subscriptions[name] = callback
}

// UnsubscribeState removes the callback set in SubscribeState
func UnsubscribeState(name string) {
	delete(subscriptions, name)
}

// Dispatch dispatches events to all subscribers
func Dispatch(action *messages.UnionEvent) {
	if action.Type != "" {
		state = reduce(&state, action)
		for _, callback := range subscriptions {
			callback(state)
		}
	}
}
