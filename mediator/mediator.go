package mediator

type CricketMediator interface {
	NotifyBallBowled(bowlder string, runs int, isWicket bool)
	NotifyDRSRequest(batsman string)
	DeclareDecision(decision string)
	GetMatchStatus() string
}

type CricketMatch struct {
	score      int
	wickets    int
	overs      float32
	umpire     *Umpire
	scoreboard *Scoreboard
	players    map[string]*Player
}

type Player struct {
	name     string
	mediator CricketMediator
}

type Scoreboard struct {
	mediator CricketMediator
}

type Umpire struct {
	mediator CricketMediator
}
