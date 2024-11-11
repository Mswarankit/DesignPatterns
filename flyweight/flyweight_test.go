package flyweight

import "testing"

func TestFlyweightPattern(t *testing.T) {
	factory := NewPlayerFactory()

	playerType1 := factory.GetPlayerType("Batsman", "Elite")
	player1 := NewPlayer(playerType1, "Virat", "Top Order")
	player2 := NewPlayer(playerType1, "Rohit", "Opener")

	if player1.playerType != player2.playerType {
		t.Error("Flyweight pattern failed: Players should share the same playerType")
	}

	playerType2 := factory.GetPlayerType("Bowler", "Elite")
	if playerType1 == playerType2 {
		t.Error("Different player types should not share the same object")
	}

	playerType3 := factory.GetPlayerType("Batsman", "Elite")
	if playerType1 != playerType3 {
		t.Error("Flyweight pattern failed: Same player type should be reused")
	}

	if playerType1.baseStats["batting"] <= playerType1.baseStats["bowling"] {
		t.Error("Batsman's batting stats should be higher than bowling stats")
	}
}
