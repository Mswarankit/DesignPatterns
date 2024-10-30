package adapter

import "fmt"

/*
Adapter -> The adapter pattern is Structural Patterns.

Adapter patterns allow incompatible interfaces to work together by converting one interface into another,
enabling integration of diverse systems.
*/

// Cricket Interface is target interface common interface for all cricket games
type CricketGame interface {
	Play() string
	GetScore() string
	GetOvers() int
}

// Standard cricket game (Adaptee)
type StandardCricketGame struct {
	gameType string
	score    int
	overs    int
}

// Defining standard cricket game
func NewStandardCricketGame(gameType string, overs int) *StandardCricketGame {
	return &StandardCricketGame{
		gameType: gameType,
		score:    0,
		overs:    overs,
	}
}

// adding play game
func (s *StandardCricketGame) PlayGame() string {
	return "Playing " + s.gameType + " cricket"
}

// Current Score of play
func (s *StandardCricketGame) CurrentScore() string {
	return fmt.Sprintf("%d runs", s.score)
}

// Test cricket game (different interface that needs adaption)
type TestCricketGame struct {
	innings  int
	score    int
	sessions int
}

// Define test cricket game
func NewTestCricketGame() *TestCricketGame {
	return &TestCricketGame{
		innings:  1,
		score:    0,
		sessions: 15, // 3 session in one day so 5 day is 15
	}
}

// Use Play test game string
func (t *TestCricketGame) PlayTestMatch() string {
	return "Playing Test cricket"
}

// Get the the test score
func (t *TestCricketGame) GetTestScore() string {
	return fmt.Sprintf("%d runs (Innings %d)", t.score, t.innings)
}

// Adapter for StandardCricketGame
type StandardGameAdapter struct {
	standardGame *StandardCricketGame
}

// Use adapter to handle the main standardcricketgame
func NewStandardGameAdapter(gametype string, overs int) *StandardGameAdapter {
	return &StandardGameAdapter{
		standardGame: NewStandardCricketGame(gametype, overs),
	}
}

func (s *StandardGameAdapter) Play() string {
	return s.standardGame.PlayGame()
}

func (s *StandardGameAdapter) GetScore() string {
	return s.standardGame.CurrentScore()
}

func (s *StandardGameAdapter) GetOvers() int {
	return s.standardGame.overs
}

// Use Testcricketgame for handling TestGameAdapter
type TestGameAdapter struct {
	testGame *TestCricketGame
}

func NewTestGameAdapter() *TestGameAdapter {
	return &TestGameAdapter{
		testGame: NewTestCricketGame(),
	}
}

func (t *TestGameAdapter) Play() string {
	return t.testGame.PlayTestMatch()
}

func (t *TestGameAdapter) GetScore() string {
	return t.testGame.GetTestScore()
}

func (t *TestGameAdapter) GetOvers() int {
	return 90
}
