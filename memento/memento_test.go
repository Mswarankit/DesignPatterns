package memento

import "testing"

func TestCricketGameMemento(t *testing.T) {
	game := NewCricketGame()
	caretaker := NewGameCareTaker()

	initialState := game.GetCurrentState()
	if initialState.Score != 0 || initialState.Wickets != 0 {
		t.Errorf("Intial state incorrect, got score: %d, wickets: %d", initialState.Score, initialState.Wickets)
	}

	game.UpdateScore(6, 0, 1.0)
	caretaker.SaveState(game.CreateMemento())
	game.UpdateScore(4, 1, 2.0)
	caretaker.SaveState(game.CreateMemento())
	game.UpdateScore(10, 0, 3.0)
	caretaker.SaveState(game.CreateMemento())

	currrentState := game.GetCurrentState()
	if currrentState.Score != 20 || currrentState.Wickets != 1 {
		t.Errorf("Current state incorrect, got score: %d, wicket: %d", currrentState.Score, currrentState.Wickets)
	}

	firstOverState, err := caretaker.RestoreState(0)
	if err != nil {
		t.Errorf("Error restoring state: %v", err)
	}
	game.RestoreFromMemento(firstOverState)
	restoredState := game.GetCurrentState()
	if restoredState.Score != 6 || restoredState.Wickets != 0 {
		t.Errorf("Restored state incorrect, got score: %d, wickets: %d", restoredState.Score, restoredState.Wickets)
	}
}

func TestInvalidStateRestore(t *testing.T) {
	game := NewCricketGame()
	_ = game.GetCurrentState()
	caretaker := NewGameCareTaker()

	_, err := caretaker.RestoreState(0)
	if err == nil {
		t.Error("Expected error when restoring invalid state")
	}
}
