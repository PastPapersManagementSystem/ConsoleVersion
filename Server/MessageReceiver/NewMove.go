package MessageReceiver

import (
	"encoding/json"
	"log"
	"main/Server/MessageReceiver/MessageTypes"
)

func NewMove(receivedMessage string) (string, error) {
	response := "success"

	var addNewPlayerToLobby MessageTypes.NewMove
	if err := json.Unmarshal([]byte(receivedMessage), &addNewPlayerToLobby); err != nil {
		log.Fatal("failed to unmarshall add new player message...\nreport: ", err)
		return failure, err
	}

	return response, nil
}
