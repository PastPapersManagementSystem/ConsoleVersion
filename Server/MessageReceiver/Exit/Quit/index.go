package Quit

import (
	"errors"
	"main/Server/MessageReceiver/Global"
	"main/Server/MessageReceiver/MessageTypes"
)

func DoesLobbyExist(lobbyId int) bool {
	for _, currentLobby := range Global.Sessions {
		if currentLobby.LobbyId == lobbyId {
			return true
		}
	}

	return false
}

func IsPlayerInGame(quit MessageTypes.Quit) bool {
	for _, currentSession := range Global.Sessions {
		if currentSession.DoesPlayerExist(quit.PlayerId) {
			return true
		}
	}

	return true
}

func RemovePlayerFromGameMembers(playerId int) {
	for _, currentSession := range Global.Sessions {
		if currentSession.DoesPlayerExist(playerId) {
			for index, member := range currentSession.Members {
				if member == playerId {
					currentSession.Members = append(currentSession.Members[:index], currentSession.Members[index+1:]...)
				}
			}
		}
	}
}

func TerminateGame(lobbyId int) {
	for index, currentSession := range Global.Sessions {
		if currentSession.LobbyId == lobbyId {
			Global.Sessions = append(Global.Sessions[:index], Global.Sessions[index+1:]...)
		}
	}
}

func IsRequesterTheCreator(quit MessageTypes.Quit) bool {
	for _, currentSession := range Global.Sessions {
		if currentSession.IsRequesterTheCreator(quit.PlayerId) {
			return true
		}
	}

	return false
}

func RemovePlayer(quit MessageTypes.Quit) (string, error) {
	if !DoesLobbyExist(quit.LobbyId) {
		return "", errors.New(lobbyDoesNotExist)
	}

	if !IsPlayerInGame(quit) {
		return "", errors.New(playerNotInGame)
	}

	if !IsRequesterTheCreator(quit) {
		RemovePlayerFromGameMembers(quit.PlayerId)
		return ok, nil
	}

	TerminateGame(quit.LobbyId)
	return ok, nil
}
