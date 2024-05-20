package Commands

import (
	"encoding/json"
	"log"
	"main/Server/MessageReceiver/MessageTypes"
	"main/Server/MessageReceiver/StartNewGame/GameLobby/AddPlayer"
)

func Add(receivedMessage string) (string, error) {
	var addNewPlayerToLobby MessageTypes.AddNewPlayer
	if err := json.Unmarshal([]byte(receivedMessage), &addNewPlayerToLobby); err != nil {
		log.Fatal("failed to unmarshall add new player message...\nreport: ", err)
		return failure, err
	}

	return AddPlayer.AddPlayer(addNewPlayerToLobby)
}
