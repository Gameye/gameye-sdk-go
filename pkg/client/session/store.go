package session

import (
	"fmt"

	messages "github.com/Gameye/sdk-messages-go/pkg/event"
)

type OnData func(State)

var state State
var subscriptions map[string]OnData

func SubscribeState(callback OnData) {
	if subscriptions == nil {
		subscriptions = make(map[string]OnData)
	}

	subscriptions[fmt.Sprintf("%d", callback)] = callback
}

func UnsubscribeState(callback OnData) {
	delete(subscriptions, fmt.Sprintf("%d", callback))
}

func Dispatch(action *messages.UnionEvent) {
	if action.Type != "" {
		state = Reduce(state, action)
		for _, callback := range subscriptions {
			callback(state)
		}
	}
}
