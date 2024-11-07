package flyweight

import "testing"

func TestFlyweightPattern(t *testing.T) {
	factory := NewPlayerFactory()

	// Test case 1: Create multiple players of the same type
	playerType1 := factory.GetPlayerType("Batsman", "Elite")
	player1 := NewPlayer(playerType1, "Virat", "Top Order")
	player2 := NewPlayer(playerType1, "Rohit", "Opener")

	// Verify that both players share the same playerType object
	if player1.playerType != player2.playerType {
		t.Error("Flyweight pattern failed: Players should share the same playerType")
	}

	// Test case 2: Verify different player types are created
	playerType2 := factory.GetPlayerType("Bowler", "Elite")
	if playerType1 == playerType2 {
		t.Error("Different player types should not share the same object")
	}

	// Test case 3: Verify flyweight reuse
	playerType3 := factory.GetPlayerType("Batsman", "Elite")
	if playerType1 != playerType3 {
		t.Error("Flyweight pattern failed: Same player type should be reused")
	}

	// Test case 4: Verify stats are correctly set
	if playerType1.baseStats["batting"] <= playerType1.baseStats["bowling"] {
		t.Error("Batsman's batting stats should be higher than bowling stats")
	}
}
