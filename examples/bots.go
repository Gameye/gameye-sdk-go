package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"../pkg/client"
	"../pkg/client/logs"
	"../pkg/client/session"
	"github.com/google/uuid"
)

func main() {
	config := client.GameyeClientConfig{
		Endpoint: "",
		Token:    "",
	}
	gameyeClient := client.NewGameyeClient(config)

	sessionID := uuid.New().String()

	log.Printf("Subscribing to session events %v\n", sessionID)
	onSessionState := func(state session.State) {
		foundSession := session.SelectSession(state, sessionID)
		if foundSession.Id != "" {
			log.Printf("Match Ready! %v", foundSession)
		}
	}

	var err error
	err = client.SubscribeSessionEvents(gameyeClient, onSessionState)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting Match %v\n", sessionID)
	client.StartMatch(
		gameyeClient,
		sessionID,
		"csgo-dem",
		[]string{"frankfurt"},
		"bots",
		map[string]interface{}{"maxRounds": 2},
		"",
	)

	currentLine := 0
	allLogs := []logs.LogLine{}
	onLogState := func(state logs.State) {
		newLogs := logs.SelectLogsSince(state, currentLine)
		for _, v := range newLogs {
			log.Printf("%d: %s", v.LineKey, v.Payload)
		}
		currentLine += len(newLogs)
		allLogs = logs.SelectAllLogs(state)
	}

	err = client.SubscribeLogEvents(gameyeClient, sessionID, onLogState)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(15 * time.Second)

	log.Printf("Stopping Match %v\n", sessionID)
	client.StopMatch(gameyeClient, sessionID)

	time.Sleep(5 * time.Second)

	session.UnsubscribeState(onSessionState)
	logs.UnsubscribeState(onLogState)

	file, err := os.Create("logs.txt")
	for _, v := range allLogs {
		io.WriteString(file, fmt.Sprintf("%d: %s", v.LineKey, v.Payload))
	}
	file.Close()
}
