package srp

import "testing"

func TestBattingStatistics(t *testing.T) {
	player := &Player{
		ID:         1,
		Name:       "Virat Kohli",
		BattingAvg: 69.08,
		Role:       "Batsman",
	}
	stats := &BattingStatistics{
		player: player,
		matches: []Match{
			{Runs: 100, Balls: 50, Result: "Won"},
			{Runs: 80, Balls: 40, Result: "Lost"},
		},
	}
	expected := 69.08
	if got := stats.CalculateBattingAverage(); got != expected {
		t.Errorf("CalculateBattingAverage() = %v, want %v", got, expected)
	}
}

func TestPlayerDatabase(t *testing.T) {
	db := &PlayerDatabase{
		players: make(map[int]Player),
	}

	player := Player{
		ID:         1,
		Name:       "Rohit Sharma",
		BattingAvg: 48.96,
		Role:       "Opening Batsman",
	}

	if err := db.SavePlayer(player); err != nil {
		t.Errorf("SavePlayer() error = %v", err)
	}
	got, err := db.GetPlayer(1)
	if err != nil {
		t.Errorf("GetPlayer() error = %v", err)
	}

	if got.Name != player.Name {
		t.Errorf("GetPlayer() = %v, want %v", got.Name, player.Name)
	}
}
