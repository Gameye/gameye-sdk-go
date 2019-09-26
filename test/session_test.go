package test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"../pkg/client/session"
)

func createSession(gameKey string, sessionID string) (string, session.Session) {
	if sessionID == "" {
		sessionID = uuid.New().String()
	}

	if gameKey == "" {
		gameKey = uuid.New().String()
	}

	created := session.Session{
		ID:       sessionID,
		Image:    gameKey,
		Location: "amsterdam",
		Host:     "127.0.0.1",
		Created:  time.Now().Unix(),
		Port:     map[string]int64{"game": 1234},
	}

	return sessionID, created
}

func TestSelectsAllSessions(test *testing.T) {
	unnamedSessions := 15
	sessions := make(map[string]session.Session, unnamedSessions)
	for i := 0; i < unnamedSessions; i++ {
		id, created := createSession("", "")
		sessions[id] = created
	}

	state := session.StateWithSessions(sessions)
	filtered := session.SelectSessionList(state)
	assert.Equal(test, unnamedSessions, len(filtered))
}

func TestSelectsOnlyRequestedSessions(test *testing.T) {
	unnamedSessions := 5
	namedSessions := 5
	sessions := make(map[string]session.Session, unnamedSessions+namedSessions)
	for i := 0; i < unnamedSessions; i++ {
		id, created := createSession("", "")
		sessions[id] = created
	}

	for i := 0; i < namedSessions; i++ {
		id, created := createSession("game-two", "")
		sessions[id] = created
	}

	state := session.StateWithSessions(sessions)
	filtered := session.SelectSessionListForGame(state, "game-two")
	assert.Equal(test, namedSessions, len(filtered))
}

func TestSelectsASessions(test *testing.T) {
	unnamedSessions := 15
	sessions := make(map[string]session.Session, unnamedSessions+1)
	for i := 0; i < unnamedSessions; i++ {
		id, created := createSession("", "")
		sessions[id] = created
	}

	id, specificSession := createSession("specific-game", "session-id-one")
	sessions[id] = specificSession

	state := session.StateWithSessions(sessions)
	foundSession := session.SelectSession(state, "session-id-one")
	assert.Equal(test, "specific-game", foundSession.Image)

	foundSession = session.SelectSession(state, "some-key-that-shouldnt-exist")
	assert.Equal(test, "", foundSession.ID)
}
