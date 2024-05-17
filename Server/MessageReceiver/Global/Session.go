package Global

type State struct {
	CurrentCall    string //card that is called.
	CurrentCaller  int    //player who called.
	NextMoveCaller int    //player who will call the next move.
	Cards          map[string][]int
}

type Session struct {
	Creator         int
	NoOfMembers     int
	Members         []string
	GameNo          int
	PreviousWinners []string
	CurrentState    State
	PreviousStates  []State
}

var Sessions []Session
