package prototype

import "testing"

func TestCricketGamePrototype(t *testing.T) {
	originalGame := NewCricketGame(100, 2, 20, []string{"Virat Kohli |", "Rohit Sharma|"})
	// Cline the original game
	clonedgame := originalGame.Clone().(*CricketGame)
	originalGame.Display()
	// Test initial clone equality
	if originalGame.Score != clonedgame.Score {
		t.Errorf("Expected cloned score to be %d, got %d", originalGame.Score, clonedgame.Score)
	}
	clonedgame.Score = 150
	clonedgame.Wickets = 4
	clonedgame.Overs = 25
	clonedgame.BatsmanNames = append(clonedgame.BatsmanNames, "Rishab Pant|")
	clonedgame.Display()

	if originalGame.Score != 100 {
		t.Errorf("Expected original score to be 100, got %d", originalGame.Score)
	}
	if len(clonedgame.BatsmanNames) != 3 {
		t.Errorf("Expected cloined batsman to be 3, got %d", len(clonedgame.BatsmanNames))
	}

}

func TestMultipleClones(t *testing.T) {
	// Create template game
	templateGame := NewCricketGame(0, 0, 0, []string{})

	// Create multiple games from template
	t20Game := templateGame.Clone().(*CricketGame)
	t20Game.Overs = 20

	testGame := templateGame.Clone().(*CricketGame)
	testGame.Overs = 90

	if t20Game.Overs != 20 {
		t.Errorf("Expected T20 game overs to be 20, got %d", t20Game.Overs)
	}
	if testGame.Overs != 90 {
		t.Errorf("Expected Test game overs to be 90, got %d", testGame.Overs)
	}
}
