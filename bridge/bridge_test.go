package bridge

import (
	"strings"
	"testing"
)

func TestCricketGameBridge(t *testing.T) {
	scoring := NewStandardScoringSystem()
	testMatch := NewTestMatch(scoring)

	if status := testMatch.GetGameStatus(); !strings.Contains(status, "0/0") {
		t.Errorf("Expected initial score 0/0, got %s", status)
	}
	testMatch.AddScore(4)
	testMatch.TakeWicket()

	runs, wickets := scoring.GetCurrentState()
	if runs != 4 || wickets != 1 {
		t.Errorf("Expected score 4/1, got %d/%d", runs, wickets)
	}

}

func TestDLSMethod(t *testing.T) {
	scoring := NewDLScoringSystem()
	odiMatch := NewODIMatch(scoring)

	odiMatch.AddScore(100)
	odiMatch.TakeWicket()

	scoring.SetRainFactor(1.2)

	status := odiMatch.GetGameStatus()
	if !strings.Contains(status, "DLS: 120.0") {
		t.Errorf("Expected DLS adjusted score around 120, got %s", status)
	}
}

func TestScoreAccumalation(t *testing.T) {
	scoring := NewStandardScoringSystem()
	game := NewT20Match(scoring)

	scores := []int{4, 6, 1, 2}
	expectedTotal := 13

	for _, score := range scores {
		game.AddScore(score)
	}

	runs, _ := scoring.GetCurrentState()
	if runs != expectedTotal {
		t.Errorf("Expected total score %d, got %d", expectedTotal, runs)
	}
}
