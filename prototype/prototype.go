package prototype

import "fmt"

type GamePrototype interface {
	Clone() GamePrototype
}

type CricketGame struct {
	Score        int
	Wickets      int
	Overs        int
	BatsmanNames []string
}

func (g *CricketGame) Clone() GamePrototype {
	newGame := *g
	newGame.BatsmanNames = make([]string, len(g.BatsmanNames))
	copy(newGame.BatsmanNames, g.BatsmanNames)
	return &newGame
}

func NewCricketGame(score, wickets, overs int, batsmanNames []string) *CricketGame {
	return &CricketGame{
		Score:        score,
		Wickets:      wickets,
		Overs:        overs,
		BatsmanNames: batsmanNames,
	}
}

func (g *CricketGame) Display() {
	fmt.Printf("Score: %d, Wickets: %d, Overs: %d, Batsmen: %v\n", g.Score, g.Wickets, g.Overs, g.BatsmanNames)
}
