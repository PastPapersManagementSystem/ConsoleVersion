package Commands

import (
	"encoding/json"
	"log"
	"main/Server/MessageReceiver/GameLobby"
	"main/Server/MessageReceiver/MessageTypes"
)

func Create(receivedMessage string) (string, error) {
	var addNewPlayerToLobby MessageTypes.CreateLobby
	if err := json.Unmarshal([]byte(receivedMessage), &addNewPlayerToLobby); err != nil {
		log.Fatal("failed to unmarshall add new player message...\nreport: ", err)
		return failure, err
	}

	return GameLobby.Create(addNewPlayerToLobby)
}
