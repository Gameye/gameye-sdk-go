package main

import (
	"log"
	"time"

	"../pkg/client"
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
	err := client.SubscribeSessionEvents(gameyeClient)
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

	time.Sleep(10 * time.Second)

	log.Printf("Stopping Match %v\n", sessionID)
	client.StopMatch(gameyeClient, sessionID)

	time.Sleep(5 * time.Second)
}
