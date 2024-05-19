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

func IsInitiatorCreator(creatorId int, currentSession *Global.Session) bool {
	if currentSession == nil {
		return false
	}

	return creatorId == currentSession.CreatorId
}

func ValidateMembersCount(session *Global.Session) bool {
	return session.CurrentMembersCount() == session.NoOfMembers
}

func StartGame(initiator MessageTypes.StartGame) (string, error) {
	currentSession := GetSessionForTheLobbyCode(initiator.LobbyCode)

	if !IsInitiatorCreator(initiator.CreatorId, currentSession) {
		return "", errors.New(UnauthorizedRequester)
	}

	if !ValidateMembersCount(currentSession) {
		return "", errors.New(IncompleteMembersCount)
	}

	currentSession.SwitchCurrentStateToStart()
	return Ok, nil
}
