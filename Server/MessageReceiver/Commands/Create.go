package Commands

import (
	"encoding/json"
	"log"
	"main/Server/MessageReceiver/MessageTypes"
	"main/Server/MessageReceiver/StartNewGame/GameLobby"
)

func Create(receivedMessage string) (string, error) {
	var createLobby MessageTypes.CreateLobby
	if err := json.Unmarshal([]byte(receivedMessage), &createLobby); err != nil {
		log.Fatal("failed to unmarshall add new player message...\nreport: ", err)
		return failure, err
	}

	return GameLobby.Create(createLobby)
}
