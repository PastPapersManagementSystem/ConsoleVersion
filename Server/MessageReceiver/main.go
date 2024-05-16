package MessageReceiver

import (
	"encoding/json"
	"log"
)

const (
	add     = "ADD"
	create  = "CREATE"
	move    = "MOVE"
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
		return Add(message)
	case create:
		return Create(message)
	case move:
		return NewMove(message)
	case unknown:
		return unknown, nil
	}

	return response, nil
}

func MessageReceiverHandler(message string) (string, error) {
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