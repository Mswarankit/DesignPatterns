package adapter

import (
	"strings"
	"testing"
)

func TestCricketGameAdapters(t *testing.T) {
	games := []CricketGame{
		NewStandardGameAdapter("T20", 20),
		NewStandardGameAdapter("ODI", 50),
		NewTestGameAdapter(),
	}
	testCases := []struct {
		gameIndex   int
		expectPlay  string
		expectOvers int
	}{
		{0, "Playing T20", 20},
		{1, "Playing ODI", 50},
		{2, "Playing Test", 90},
	}

	for _, tc := range testCases {
		game := games[tc.gameIndex]

		playResult := game.Play()
		if !strings.Contains(playResult, tc.expectPlay) {
			t.Errorf("Expected play result to contain %s, got %s", tc.expectPlay, playResult)
		}
		score := game.GetScore()
		if score == "" {
			t.Error("Expected non-empty score string")
		}
	}
}

func TestSpecificGameBehavior(t *testing.T) {
	// Test Standard Game Adapter
	t20Game := NewStandardGameAdapter("T20", 20)
	if t20Game.GetOvers() != 20 {
		t.Errorf("Expected T20 game to have 20 overs, got %d",
			t20Game.GetOvers())
	}

	// Test Test Game Adapter
	testGame := NewTestGameAdapter()
	playResult := testGame.Play()
	if !strings.Contains(playResult, "Test cricket") {
		t.Errorf("Expected Test cricket in play result, got %s", playResult)
	}
}
