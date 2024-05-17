package MessageReceiver

import (
	"encoding/json"
	"log"
	"main/Server/MessageReceiver/Commands"
)

const (
	add     = "ADD"
	create  = "CREATE"
	move    = "MOVE"
	end     = "END"
	quit    = "QUIT"
	unknown = "UNKNOWN"
	failure = "FAILURE"
)

func UnMarshallReceivedMessage(message string) (map[string]interface{}, error) {
	var receivedMessage map[string]interface{}
	messageUnmarshalError := json.Unmarshal([]byte(message), &receivedMessage)

	return receivedMessage, messageUnmarshalError
}

func HandleReceivedMessage(message string) (string, error) {
	response := "smooth"

	switch message {
	case add:
		return Commands.Add(message)
	case create:
		return Commands.Create(message)
	case move:
		return Commands.NewMove(message)
	case unknown:
		return unknown, nil
	}

	return response, nil
}

func ReceiverHandler(message string) (string, error) {
	receivedMessage, unMarshallError := UnMarshallReceivedMessage(message)
	if unMarshallError != nil {
		log.Fatal("failed to unmarshall received message\nreport: ", unMarshallError)
		return "", unMarshallError
	}

	response, responseError := HandleReceivedMessage(receivedMessage["type"].(string))
	if responseError != nil {
		log.Fatal("failed to handle received message\nreport: ", responseError)
		return "", responseError
	}

	return response, nil
}
