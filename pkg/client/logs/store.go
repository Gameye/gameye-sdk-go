package logs

import (
	"fmt"

	"../patch"
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

func Dispatch(patches *[]patch.Patch) {
	state = Reduce(&state, patches)
	for _, callback := range subscriptions {
		callback(state)
	}
}
