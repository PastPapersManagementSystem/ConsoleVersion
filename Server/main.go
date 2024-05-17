package main

import (
	"github.com/pebbe/zmq4"
	"log"
	"main/Server/MessageReceiver"
)

func main() {
	zeroContext, zeroContextError := zmq4.NewContext()

	if zeroContextError != nil {
		log.Fatal("failed to start server due to server crashing...\nreport: ", zeroContextError)
		return
	}

	server, serverError := zeroContext.NewSocket(zmq4.REP)
	if serverError != nil {
		log.Fatal("failed to start server due to server crashing...\nreport: ", serverError)
		return
	}

	for {
		message, messageError := server.Recv(0)
		if messageError != nil {
			log.Fatal("failed to retrieve message...\nreport: ", messageError)
			return
		}

		log.Printf("Received message: %s\n", message)

		response, messageReceiverHandlerError := MessageReceiver.ReceiverHandler(message)

		if messageReceiverHandlerError != nil {
			log.Fatal("failed to get receive message\nreport: ", messageReceiverHandlerError)
			return
		}

		send, messageSendError := server.Send(response, 0)

		if messageSendError != nil {
			log.Fatal("failed to send message...\nreport: ", messageSendError)
			return
		}

		log.Printf("Send message: %d\n", send)
	}
}
