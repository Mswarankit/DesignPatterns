package prototype

import "fmt"

/*
The prototype patterns is a creational patterns.
Allows an object to clone itself, creating new instances based on existing object without needing
to create details of how to create them.
*/

// GamePrototype interface defines the method to clone game instances
type GamePrototype interface {
	Clone() GamePrototype
}

// CricketGame represents the prototype that can be cloned
type CricketGame struct {
	Score        int
	Wickets      int
	Overs        int
	BatsmanNames []string
}

// Clone create a copy of the  current CricketGame
// Create a deep copy of the CricketGame
// Create a new Slice for Batsman to avoid sharing the reference
func (g *CricketGame) Clone() GamePrototype {
	newGame := *g
	newGame.BatsmanNames = make([]string, len(g.BatsmanNames))
	copy(newGame.BatsmanNames, g.BatsmanNames)
	return &newGame
}

// NewCricketGame creates a new instance of CricketGame
func NewCricketGame(score, wickets, overs int, batsmanNames []string) *CricketGame {
	return &CricketGame{
		Score:        score,
		Wickets:      wickets,
		Overs:        overs,
		BatsmanNames: batsmanNames,
	}
}

// Display the current Game status
func (g *CricketGame) Display() {
	fmt.Printf("Score: %d, Wickets: %d, Overs: %d, Batsmen: %v\n", g.Score, g.Wickets, g.Overs, g.BatsmanNames)
}
