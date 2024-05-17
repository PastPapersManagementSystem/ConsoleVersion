package GameLobby

import (
	"errors"
	"main/Server/MessageReceiver/Global"
	"main/Server/MessageReceiver/MessageTypes"
)

func ValidateCreator(createId int) bool {

	for _, currentState := range Global.Sessions {
		if currentState.CreatorId == createId {
			return false
		}
	}

	return true
}

func InitializeNewSession(createNewLobby MessageTypes.CreateLobby) {
	var newSession Global.Session
	newSession.InitializeNewSession(createNewLobby.CreatorId, createNewLobby.MaxMembersInLobby)
}

func Create(createNewLobby MessageTypes.CreateLobby) (string, error) {
	canCreatorCreateLobby := ValidateCreator(createNewLobby.CreatorId)

	if canCreatorCreateLobby {
		return "", errors.New(cantCreateLobby)
	}

	InitializeNewSession(createNewLobby)
	return success, nil
}
