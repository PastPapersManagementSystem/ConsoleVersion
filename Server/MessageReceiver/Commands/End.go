package Commands

import (
	"encoding/json"
	"log"
	EndGame2 "main/Server/MessageReceiver/EndGame"
	"main/Server/MessageReceiver/MessageTypes"
)

func EndGame(message string) (string, error) {
	var endGame MessageTypes.EndGame
	if err := json.Unmarshal([]byte(message), &endGame); err != nil {
		log.Fatal("failed to unmarshall end game message...\nreport: ", err)
		return failure, err
	}

	return EndGame2.TerminateGame(endGame)
}
