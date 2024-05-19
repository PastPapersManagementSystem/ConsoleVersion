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
	LobbyId            int      //randomly generated lobby code
	CreatorId          int      //the id whose created the game
	NoOfMembers        int      //no of members required to start the game
	CurrentNoOfMembers int      //the present number of members
	CurrentState       string   //tells about the state of the session, either in waiting or completed
	Members            []int    //the members alongside their id
	GameNo             int      //what is the game number in a row they are playing
	PreviousWinners    []string //name of the players who won previously
	CurrentGameState   State    //what is the current state of the game, has all the cards information and etc
	PreviousStates     []State  //how the game progressed
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
