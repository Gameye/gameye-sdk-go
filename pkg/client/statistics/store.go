package statistics

import (
	"../patch"
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

func Dispatch(patches *[]patch.Patch) {
	state = reduce(&state, patches)
	for _, callback := range subscriptions {
		callback(state)
	}
}
