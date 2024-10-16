package enums

type GameType string

const (
	Quick     GameType = "quick"
	Arena     GameType = "arena"
	Arena2    GameType = "arena_2_0"
	Ranked    GameType = "ranked"
	Ranked2   GameType = "ranked_2_0"
	ProLadder GameType = "pro_ladder"
)

type GameEnd string

const (
	Win        GameEnd = "Win"
	Draw       GameEnd = "Draw"
	NotStarted GameEnd = "NotStarted"
)

type PlayerStatus string

const (
	Ok            PlayerStatus = "Ok"
	Forfeit       PlayerStatus = "Forfeit"
	Disconnected  PlayerStatus = "Disconnected"
	NotConnected  PlayerStatus = "NotConnected"
	NotResponding PlayerStatus = "NotResponding"
)
