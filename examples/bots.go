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
	"../pkg/client/statistics"
	"github.com/google/uuid"
)

func main() {
	var err error

	config := client.GameyeClientConfig{
		Endpoint: "",
		Token:    "",
	}

	gameyeClient, err := client.NewGameyeClient(config)
	handleErr(err)

	sessionID := uuid.New().String()

	log.Printf("Subscribing to session events %v\n", sessionID)
	onSessionState := func(state session.State) {
		foundSession := session.SelectSession(state, sessionID)
		if foundSession.Id != "" {
			log.Printf("Match Ready! %v", foundSession)
		}
	}

	err = client.SubscribeSessionEvents(gameyeClient, onSessionState)
	handleErr(err)

	log.Printf("Starting Match %v\n", sessionID)
	err = client.StartMatch(
		gameyeClient,
		sessionID,
		"csgo-dem",
		[]string{"frankfurt"},
		"bots",
		map[string]interface{}{"maxRounds": 2},
		"",
	)
	handleErr(err)

	currentLine := 0
	var allLogs []logs.LogLine
	onLogState := func(state logs.State) {
		newLogs := logs.SelectLogsSince(state, currentLine)
		for _, v := range newLogs {
			log.Printf("%d: %s", v.LineKey, v.Payload)
		}
		currentLine += len(newLogs)
		allLogs = logs.SelectAllLogs(state)
	}

	err = client.SubscribeLogEvents(gameyeClient, sessionID, onLogState)
	handleErr(err)

	rawStats := ""
	onStatisticsState := func(state statistics.State) {
		statistics.SelectPlayerList(state)
		rawStats, err = statistics.SelectRawStatistics(state)
		handleErr(err)
	}

	err = client.SubscribeStatisticsEvents(gameyeClient, sessionID, onStatisticsState)
	handleErr(err)

	time.Sleep(30 * time.Second)

	log.Printf("Stopping Match %v\n", sessionID)
	err = client.StopMatch(gameyeClient, sessionID)
	handleErr(err)

	time.Sleep(5 * time.Second)

	file, err := os.Create("logs.txt")
	if file != nil && allLogs != nil {
		for _, v := range allLogs {
			_, err = io.WriteString(file, fmt.Sprintf("%d: %s", v.LineKey, v.Payload))
		}

		err = file.Close()
		handleErr(err)
	}

	file, err = os.Create("stats.txt")
	if file != nil {
		_, err = io.WriteString(file, rawStats)
		err = file.Close()
		handleErr(err)
	}

}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
