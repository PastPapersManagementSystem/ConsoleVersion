package Global

import "math/rand"

type State struct {
	CurrentCall    string //card that is called.
	CurrentCaller  int    //player who called.
	NextMoveCaller int    //player who will call the next move.
	Cards          map[string][]int
}

const (
	waiting   = "WAITING"
	completed = "COMPLETED"
)

type Session struct {
	LobbyId            int
	CreatorId          int
	NoOfMembers        int
	CurrentNoOfMembers int
	CurrentState       string
	Members            []int
	GameNo             int
	PreviousWinners    []string
	CurrentGameState   State
	PreviousStates     []State
}

func GenerateRandomLobbyId() int {
	minimumLobbyCode := 100000
	maximumLobbyCode := 999999

	randomNumber := rand.Intn(maximumLobbyCode-minimumLobbyCode) + minimumLobbyCode

	return randomNumber
}

func (session Session) InitializeNewSession(creatorId int, noOfMembers int) int {
	session.LobbyId = GenerateRandomLobbyId()
	session.CreatorId = creatorId
	session.NoOfMembers = noOfMembers
	session.Members = make([]int, noOfMembers)
	session.PreviousStates = make([]State, noOfMembers)
	session.Members = append(session.Members, creatorId)
	session.CurrentState = waiting

	return session.LobbyId
}

func (session Session) GetCurrentState() string {
	if len(session.Members) == session.NoOfMembers {
		session.CurrentState = completed
		return session.CurrentState
	}

	return session.CurrentState
}

func (session Session) IsRequesterTheCreator(requester int) bool {
	return requester == session.CreatorId
}

func (session Session) CurrentMembersCount() int {
	return len(session.Members)
}

var Sessions []Session
