package bridge

import "fmt"

type ScoreSystem interface {
	GetScore() string
	AddRuns(runs int)
	AddWicket()
	GetCurrentState() (int, int)
}

type StatndardScoringSystem struct {
	runs    int
	wickets int
}

func NewStandardScoringSystem() *StatndardScoringSystem {
	return &StatndardScoringSystem{}
}

func (s *StatndardScoringSystem) GetScore() string {
	return fmt.Sprintf("%d/%d", s.runs, s.wickets)
}

func (s *StatndardScoringSystem) AddRuns(runs int) {
	s.runs += runs
}

func (s *StatndardScoringSystem) AddWicket() {
	s.wickets++
}

func (s *StatndardScoringSystem) GetCurrentState() (int, int) {
	return s.runs, s.wickets
}

type Game interface {
	Play()
	AddScore(runs int)
	TakeWicket()
	GetGameStatus() string
}

type BaseCricketGame struct {
	scoringSystem ScoreSystem
	format        string
	overs         int
}

func (g *BaseCricketGame) Play() {
	fmt.Printf("Playing %s Cricket Match\n", g.format)
}

func (g *BaseCricketGame) AddScore(runs int) {
	g.scoringSystem.AddRuns(runs)
}

func (g *BaseCricketGame) TakeWicket() {
	g.scoringSystem.AddWicket()
}

func (g *BaseCricketGame) GetGameStatus() string {
	return fmt.Sprintf("%s Match - %s", g.format, g.scoringSystem.GetScore())
}

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

type DLSScoringSystem struct {
	StatndardScoringSystem
	rainFactor float64
}

func NewDLScoringSystem() *DLSScoringSystem {
	return &DLSScoringSystem{
		rainFactor: 1.0,
	}
}

func (d *DLSScoringSystem) GetScore() string {
	adjustedScore := float64(d.runs) * d.rainFactor
	return fmt.Sprintf("%d/%d (DLS: %.1f)", d.runs, d.wickets, adjustedScore)
}

func (d *DLSScoringSystem) SetRainFactor(factor float64) {
	d.rainFactor = factor
}
