package memento

import "fmt"

/*
Memenoto is behavioral design patterns that allows you to capture and save the internal
state of an object without exposing its details.
The Memento patterns allows us to save and restore the state of a cricket game (like saving match
situation for replay or undo functionality) without exposing
*/

// CricketGameState represents the state of a cricket match
type CricketGameState struct {
	Score        int
	Wickets      int
	CurrentOver  float32
	BatsmanScore int
	BowlerStats  string
}

// CricketGame is the originator that maintains the game state
type CricketGame struct {
	state CricketGameState
}

// NewCricketGame creates a new cricket game
func NewCricketGame() *CricketGame {
	return &CricketGame{
		state: CricketGameState{
			Score:        0,
			Wickets:      0,
			CurrentOver:  0.0,
			BatsmanScore: 0,
			BowlerStats:  "0/0",
		},
	}
}

// UpdateScore updates the game state
func (g *CricketGame) UpdateScore(runs int, wickets int, overs float32) {
	g.state.Score += runs
	g.state.Wickets += wickets
	g.state.CurrentOver = overs
	g.state.BatsmanScore += runs
}

// CreateMemento saves the current state
func (g *CricketGame) CreateMemento() *GameMemento {
	return &GameMemento{
		state: g.state,
	}
}

// RestoreFromMemento restores a previous state
func (g *CricketGame) RestoreFromMemento(m *GameMemento) {
	g.state = m.state
}

// GetCurrentState returns the current game state
func (g *CricketGame) GetCurrentState() CricketGameState {
	return g.state
}

// GameMemento stores the game state
type GameMemento struct {
	state CricketGameState
}

// GameCareTaker manages the saved states
type GameCareTaker struct {
	mementos []*GameMemento
}

// NewGameCaretaker creates a new character
func NewGameCareTaker() *GameCareTaker {
	return &GameCareTaker{
		mementos: make([]*GameMemento, 0),
	}
}

// SaveState adds a new memento to the history
func (c *GameCareTaker) SaveState(m *GameMemento) {
	c.mementos = append(c.mementos, m)
}

// RestoreState retrieves a specific state
func (c *GameCareTaker) RestoreState(index int) (*GameMemento, error) {
	if index < 0 || index >= len(c.mementos) {
		return nil, fmt.Errorf("Invalid state index")
	}
	return c.mementos[index], nil
}
