package Commands

import (
	"encoding/json"
	"log"
	"main/Server/MessageReceiver/MessageTypes"
	"main/Server/MessageReceiver/StartNewGame/StartGame"
)

func InitiateGame(receivedMessage string) (string, error) {
	var startGame MessageTypes.StartGame
	if err := json.Unmarshal([]byte(receivedMessage), &startGame); err != nil {
		log.Fatal("failed to unmarshal initiate game message...\nreport: ", err)
		return failure, err
	}

	return StartGame.StartGame(startGame)
}
