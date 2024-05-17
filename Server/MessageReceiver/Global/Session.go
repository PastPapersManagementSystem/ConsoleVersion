package Global

type State struct {
	CurrentCall    string //card that is called.
	CurrentCaller  int    //player who called.
	NextMoveCaller int    //player who will call the next move.
	Cards          map[string][]int
}

type Session struct {
	SessionId          int
	CreatorId          int
	NoOfMembers        int
	CurrentNoOfMembers int
	Members            []int
	GameNo             int
	PreviousWinners    []string
	CurrentState       State
	PreviousStates     []State
}

func (session Session) InitializeNewSession(creatorId int, noOfMembers int) {
	session.CreatorId = creatorId
	session.NoOfMembers = noOfMembers
	session.Members = make([]int, noOfMembers)
	session.PreviousStates = make([]State, noOfMembers)
	session.Members = append(session.Members, creatorId)
}

func (session Session) CurrentMembersCount() int {
	return len(session.Members)
}

var Sessions []Session
