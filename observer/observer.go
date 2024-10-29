package observer

import "fmt"

/*
Observer pattern is behavioral design patterns.
One-to-many dependency between objects, allowing multiple
observers to be notified changes in subjects state
Subject -> CricketGame
Observer -> Player and commentrator
*/

// Subject interface defines methods for managing observers
type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObserver()
}

// Observer interface defines the update method for observers
type Observer interface {
	Update(Score int)
}

// CricketGame acts as the Subject
type CricketGame struct {
	observers []Observer
	score     int
}

func NewCricketGame() *CricketGame {
	return &CricketGame{
		observers: []Observer{},
		score:     0,
	}
}

// Register Observer with as many observer required
func (g *CricketGame) RegisterObserver(o Observer) {
	g.observers = append(g.observers, o)
}

// Remove Observer one by one
func (g *CricketGame) RemoveObserver(o Observer) {
	for one, value := range g.observers {
		if value == o {
			g.observers = append(g.observers[:one], g.observers[one+1:]...)
			break
		}
	}
}

// Notify Observer with latest updates
func (g *CricketGame) NotifyObserver() {
	for _, value := range g.observers {
		value.Update(g.score)
	}
}

// Update score of the game
func (g *CricketGame) UpdateScore(score int) {
	g.score += score
	g.NotifyObserver()
}

// Player implements observer patterns
type Player struct {
	Name string
}

func (p *Player) Update(score int) {
	fmt.Printf("%s received the score update: %d\n", p.Name, score)
}

// Commentator implements the Observer interface
type Commentator struct {
	Name string
}

func (c *Commentator) Update(score int) {
	fmt.Printf("%s received the score update: %d\n", c.Name, score)
}
