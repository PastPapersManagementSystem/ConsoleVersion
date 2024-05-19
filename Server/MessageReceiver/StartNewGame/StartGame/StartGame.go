package StartGame

import (
	"errors"
	"main/Server/MessageReceiver/Global"
	"main/Server/MessageReceiver/MessageTypes"
)

func GetSessionForTheLobbyCode(lobbyCode int) *Global.Session {
	for _, currentSession := range Global.Sessions {
		if currentSession.LobbyId == lobbyCode {
			return &currentSession
		}
	}

	return nil
}

func IsInitiatorCreator(creatorId int, lobbyCode int) bool {
	currentSession := GetSessionForTheLobbyCode(lobbyCode)
	if currentSession == nil {
		return false
	}

	return creatorId == currentSession.CreatorId
}

func StartGame(initiator MessageTypes.StartGame) (string, error) {
	if !IsInitiatorCreator(initiator.CreatorId, initiator.LobbyCode) {
		return "", errors.New(UnauthorizedRequester)
	}

	return "", nil
}
