package AddPlayer

import (
	"errors"
	"main/Server/MessageReceiver/Global"
	"main/Server/MessageReceiver/MessageTypes"
)

func CheckIfLobbyExists(lobbyCode int) bool {
	for _, session := range Global.Sessions {
		if session.LobbyId == lobbyCode {
			return true
		}
	}

	return false
}

func CheckIfPlayerAlreadyInAnotherLobby(playerId int) bool {
	for _, session := range Global.Sessions {
		for _, member := range session.Members {
			if member == playerId {
				return true
			}
		}
	}

	return false
}

func AppendPlayerToTheSession(playerId int, lobbyCode int) {
	for _, session := range Global.Sessions {
		if session.LobbyId == lobbyCode {
			session.AppendNewPlayer(playerId)
			return
		}
	}
}

func AddPlayer(player MessageTypes.AddNewPlayer) (string, error) {
	doesLobbyExist := CheckIfLobbyExists(player.LobbyCode)
	if !doesLobbyExist {
		return "", errors.New(lobbyDoesNotExist)
	}

	isPlayerAlreadyInAnotherLobby := CheckIfPlayerAlreadyInAnotherLobby(player.PlayerId)
	if isPlayerAlreadyInAnotherLobby {
		return "", errors.New(playerAlreadyInAnotherLobby)
	}

	AppendPlayerToTheSession(player.PlayerId, player.LobbyCode)
	return ok, nil
}
