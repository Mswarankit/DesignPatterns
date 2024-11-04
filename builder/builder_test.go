package builder

import "testing"

func TestCricketGameBuilder(t *testing.T) {
	t.Run("Test T20 Game Configuration", func(t *testing.T) {
		builder := NewCricketGameBuilder("Team India")
		game := builder.
			SetLocation("Chinnaswami Stadium").
			AddPlayers("Rohit Sharma").AddPlayers("Jasprit Bumrah").AddPlayers("Hardik Pandya").AddPlayers("Virat Kohli").
			AddUmpires("Kumar Dharamasena").AddUmpires("Richard Kettleborough").
			SetDRS(true).
			SetFloodLights(true).
			Build()

		if game.Overs != 20 {
			t.Errorf("Expected 20 Overs for T20, got %d", game.Overs)
		}

		if game.Format != "T20" {
			t.Errorf("Expected T20 format, got %s", game.Format)
		}

	})
	t.Run("Testing Test Game Configuration", func(t *testing.T) {
		builder := NewCricketGameBuilder("Team India")
		game := builder.SetOvers(90).SetLocation("Whankhade Stadium").
			AddPlayers("Rohit Sharma").AddPlayers("Rishab Pant").Build()

		if game.Overs != 90 {
			t.Errorf("Expected 90 overs for test, got %d", game.Overs)
		}
		if game.Format != "Test" {
			t.Errorf("Expected Test format, got %s", game.Format)
		}
	})

	t.Run("Test Match Officials Configuration", func(t *testing.T) {
		builder := NewCricketGameBuilder("Test Team")
		game := builder.
			AddUmpires("Umpire 1").
			AddUmpires("Umpire 2").
			SetDRS(true).
			Build()

		if len(game.Umpires) != 2 {
			t.Errorf("Expected 2 umpires, got %d", len(game.Umpires))
		}
		if !game.HasDrs {
			t.Error("Expected DRS to be enabled")
		}
	})

	t.Run("Test Test Match Configuration", func(t *testing.T) {
		builder := NewCricketGameBuilder("England")
		game := builder.
			SetOvers(90).
			SetLocation("Lords").
			SetFloodLights(false).
			Build()

		if game.Format != "Test" {
			t.Errorf("Expected Test format, got %s", game.Format)
		}
		if game.FloodLights {
			t.Error("Expected flood lights to be disabled for Test match")
		}
	})
}
