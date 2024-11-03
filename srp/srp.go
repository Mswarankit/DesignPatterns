package srp

import "fmt"

/*
SRP -> Single Responsibility Principle
A class should have only one reason to change.

*/

// Player reresents a cricket player with basic information
type Player struct {
	ID         int
	Name       string
	BattingAvg float64
	Role       string
}

// BattingStatistics handles only batting related calculations
type BattingStatistics struct {
	player  *Player
	matches []Match
}

func (bs *BattingStatistics) CalculateBattingAverage() float64 {
	return bs.player.BattingAvg
}

// PlayerDatabase handles only player data storage operation
type PlayerDatabase struct {
	players map[int]Player
}

func (pd *PlayerDatabase) SavePlayer(player Player) error {
	pd.players[player.ID] = player
	return nil
}

func (pd *PlayerDatabase) GetPlayer(id int) (Player, error) {
	player, exists := pd.players[id]
	if !exists {
		return Player{}, fmt.Errorf("player not found")
	}
	return player, nil
}

type Match struct {
	Runs   int
	Balls  int
	Result string
}
