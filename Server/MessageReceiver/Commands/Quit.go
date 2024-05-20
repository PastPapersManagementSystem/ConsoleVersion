package Commands

import (
	"encoding/json"
	"log"
	Quit2 "main/Server/MessageReceiver/Exit/Quit"
	"main/Server/MessageReceiver/MessageTypes"
)

func Quit(receivedMessage string) (string, error) {
	var quitPlayer MessageTypes.Quit
	if err := json.Unmarshal([]byte(receivedMessage), &quitPlayer); err != nil {
		log.Fatal("failed to unmarshall quit player message...\nreport: ", err)
		return failure, err
	}

	return Quit2.RemovePlayer(quitPlayer)
}
