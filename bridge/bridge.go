package bridge

import "fmt"

/*
Bridge pattern is about preferring composition over inheritance. It allows
decoupling the abstraction (game format) from its implementation (scoring system).
This pattern makes it easy to add new game formats and scoring systems independently.
*/

// ScoreSystem defines the interface that different scoring systems must implement.
type ScoreSystem interface {
	GetScore() string
	AddRuns(runs int)
	AddWicket()
	GetCurrentState() (int, int)
}

// StandardScoringSystem implements basic scoring rules
type StatndardScoringSystem struct {
	runs    int
	wickets int
}

// Intilizing Standard Scoring system
func NewStandardScoringSystem() *StatndardScoringSystem {
	return &StatndardScoringSystem{}
}

// GetScore will format runs and wickets
func (s *StatndardScoringSystem) GetScore() string {
	return fmt.Sprintf("%d/%d", s.runs, s.wickets)
}

// AddRuns will add runs of the system
func (s *StatndardScoringSystem) AddRuns(runs int) {
	s.runs += runs
}

// AddWicket increments the number of wickets by 1
func (s *StatndardScoringSystem) AddWicket() {
	s.wickets++
}

// Provide GetCurrentState() method will return int, int
func (s *StatndardScoringSystem) GetCurrentState() (int, int) {
	return s.runs, s.wickets
}

// Game is the interface representing a cricket game.
// It defines the common operations for different cricket game formats.
type Game interface {
	Play()
	AddScore(runs int)
	TakeWicket()
	GetGameStatus() string
}

// BaseCricketGame provides the common functionality for different game formats.
type BaseCricketGame struct {
	scoringSystem ScoreSystem
	format        string
	overs         int
}

// Play prints a message indicating the game format being played.
func (g *BaseCricketGame) Play() {
	fmt.Printf("Playing %s Cricket Match\n", g.format)
}

// AddScore adds runs to the current game via the scoring system.
func (g *BaseCricketGame) AddScore(runs int) {
	g.scoringSystem.AddRuns(runs)
}

// TakeWicket records a wicket in the game.
func (g *BaseCricketGame) TakeWicket() {
	g.scoringSystem.AddWicket()
}

// GetGameStatus returns the current score of the game with its format.
func (g *BaseCricketGame) GetGameStatus() string {
	return fmt.Sprintf("%s Match - %s", g.format, g.scoringSystem.GetScore())
}

// TestMatch represents a specific type of cricket match: Test format.
// Embeds BaseCricketGame for shared functionality
type TestMatch struct {
	BaseCricketGame
}

// NewTestMatch creates and returns a new TestMatch instance.
func NewTestMatch(scoringSystem ScoreSystem) *TestMatch {
	return &TestMatch{
		BaseCricketGame: BaseCricketGame{
			scoringSystem: scoringSystem,
			format:        "Test",
			overs:         90,
		},
	}
}

// ODIMatch represents a One-Day International (ODI) cricket match.
// Embeds BaseCricketGame for shared functionality
type ODIMatch struct {
	BaseCricketGame
}

// NewODIMatch creates and returns a new ODIMatch instance.
func NewODIMatch(scoringSystem ScoreSystem) *ODIMatch {
	return &ODIMatch{
		BaseCricketGame: BaseCricketGame{
			scoringSystem: scoringSystem,
			format:        "ODI",
			overs:         50,
		},
	}
}

// T20Match represents a T20 cricket match.
// Embeds BaseCricketGame for shared functionality
type T20Match struct {
	BaseCricketGame
}

// NewT20Match creates and returns a new T20Match instance.
func NewT20Match(scoringSystem ScoreSystem) *T20Match {
	return &T20Match{
		BaseCricketGame: BaseCricketGame{
			scoringSystem: scoringSystem,
			format:        "T20",
			overs:         20,
		},
	}
}

/*
DLSScoringSystem is a specialized scoring system that includes a rain factor,
used for adjusting scores during weather-affected matches.
Embeds the standard scoring system to reuse its methods
A factor to adjust the score for rain-affected games
*/
type DLSScoringSystem struct {
	StatndardScoringSystem
	rainFactor float64
}

// NewDLScoringSystem creates and returns a new DLSScoringSystem instance with a default rain factor.
func NewDLScoringSystem() *DLSScoringSystem {
	return &DLSScoringSystem{
		rainFactor: 1.0,
	}
}

// GetScore returns the adjusted score considering the rain factor.
// Adjust the runs based on the rain factor and format the score accordingly
func (d *DLSScoringSystem) GetScore() string {
	adjustedScore := float64(d.runs) * d.rainFactor
	return fmt.Sprintf("%d/%d (DLS: %.1f)", d.runs, d.wickets, adjustedScore)
}

// SetRainFactor sets the rain factor for DLS scoring system.
func (d *DLSScoringSystem) SetRainFactor(factor float64) {
	d.rainFactor = factor
}
