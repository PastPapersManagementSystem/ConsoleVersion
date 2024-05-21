package MessageTypes

type NewMove struct {
	PlayerId  int
	LobbyId   int
	MoveType  string
	MoveValue string
}
