package observer

import "testing"

func TestObserverGame(t *testing.T) {
	game := NewCricketGame()

	player := &Player{Name: "Virat Kohli"}
	commentator := &Commentator{Name: "Ravi Shastri"}

	game.RegisterObserver(player)
	game.RegisterObserver(commentator)

	game.UpdateScore(4)
	game.UpdateScore(6)

	game.RemoveObserver(player)
	game.UpdateScore(2)

}
