package session

import (
	"log"

	messages "github.com/Gameye/sdk-messages-go/pkg/event"
)

var state State

func Subscribe() {

}

func Dispatch(action *messages.UnionEvent) {
	if action.Type != "" {
		state = Reduce(state, action)
		log.Printf("Reduced state: %v", state)
	}
}
