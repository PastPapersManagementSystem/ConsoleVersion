package EndGame

import (
	"errors"
	"main/Server/MessageReceiver/Global"
	"main/Server/MessageReceiver/MessageTypes"
)

func DoesLobbyExist(lobbyId int) bool {
	for _, currentSession := range Global.Sessions {
		if currentSession.LobbyId == lobbyId {
			return true
		}
	}
	return false
}

func IsRequesterTheCreator(lobbyId, playerId int) bool {
	for _, currentSession := range Global.Sessions {
		if currentSession.LobbyId == lobbyId && currentSession.CreatorId == playerId {
			return true
		}
	}
	return false
}

func RemoveSessionFromGlobal(lobbyId int) {
	for index, currentSession := range Global.Sessions {
		if currentSession.LobbyId == lobbyId {
			Global.Sessions = append(Global.Sessions[:index], Global.Sessions[index+1:]...)
		}
	}
}

func TerminateGame(endGame MessageTypes.EndGame) (string, error) {
	if !DoesLobbyExist(endGame.LobbyCode) {
		return "", errors.New(lobbyDoesntExist)
	}

	if !IsRequesterTheCreator(endGame.LobbyCode, endGame.PlayerId) {
		return "", errors.New(onlyCreatorCanEndGame)
	}

	RemoveSessionFromGlobal(endGame.LobbyCode)
	return "", nil
}
