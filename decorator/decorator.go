package decorator

import "fmt"

// CricketGame interface defines the methods that any cricket game should implement.
type CricketGame interface {
	Play() string
	GetScore() string
}

type BasicCricketGame struct {
	score int // The current score of the game.
}

func NewBasicCricketGame() *BasicCricketGame {
	return &BasicCricketGame{
		score: 0,
	}
}

func (g *BasicCricketGame) Play() string {
	return "Playing Cricket" // Basic message indicating the game is being played.
}

func (g *BasicCricketGame) GetScore() string {
	return fmt.Sprintf("Score: %d", g.score) // Returns the current score.
}

// CricketGameDecorator struct serves as a base decorator that holds a CricketGame instance.
type CricketGameDecorator struct {
	game CricketGame // The wrapped CricketGame instance.
}

// Play method delegates the call to the wrapped game’s Play method.
func (c *CricketGameDecorator) Play() string {
	return c.game.Play() // Calls the Play method of the wrapped game.
}

// GetScore method delegates the call to the wrapped game’s GetScore method.
func (c *CricketGameDecorator) GetScore() string {
	return c.game.GetScore() // Calls the GetScore method of the wrapped game.
}

// CommentaryDecorator struct extends CricketGameDecorator to add commentary.
type CommentaryDecorator struct {
	CricketGameDecorator        // Embedding the base decorator to reuse its methods.
	commentator          string // The name of the commentator.
}

// NewCommentratorDecorator initializes a new CommentaryDecorator with the specified game and commentator.
func NewCommentratorDecorator(game CricketGame, commentator string) *CommentaryDecorator {
	return &CommentaryDecorator{
		CricketGameDecorator: CricketGameDecorator{game: game}, // Initialize the base decorator.
		commentator:          commentator,
	}
}

// Play method overrides the Play method to add commentary details.
func (co *CommentaryDecorator) Play() string {
	return co.game.Play() + fmt.Sprintf(" with commentary by %s", co.commentator) // Adds commentary info to the play message.
}

// GetScore method overrides the GetScore method to add commentary about the game.
func (co *CommentaryDecorator) GetScore() string {
	return co.game.GetScore() + fmt.Sprintf(" What a fantastic game!") // Adds a remark about the game.
}

// StatisticDecorator struct extends CricketGameDecorator to add live statistics.
type StatisticDecorator struct {
	CricketGameDecorator         // Embedding the base decorator to reuse its methods.
	runRate              float64 // The current run rate of the game.
	boundaries           int     // The number of boundaries scored.
}

// NewStatisticsDecorators initializes a new StatisticDecorator for the specified game.
func NewStatisticsDecorators(game CricketGame) *StatisticDecorator {
	return &StatisticDecorator{
		CricketGameDecorator: CricketGameDecorator{game: game}, // Initialize the base decorator.
		runRate:              0,                                // Default run rate is set to 0.
		boundaries:           0,                                // Default number of boundaries is set to 0.
	}
}

// Play method overrides the Play method to indicate live statistics are being shown.
func (s *StatisticDecorator) Play() string {
	return s.game.Play() + " with live statistics" // Adds indication of live statistics.
}

// GetScore method overrides the GetScore method to include statistics.
func (s *StatisticDecorator) GetScore() string {
	return fmt.Sprintf("%s (Run Rate: %.2f, Boundaries: %d)", s.game.GetScore(), s.runRate, s.boundaries) // Returns score with statistics.
}

// LiveStreamingDecorator struct extends CricketGameDecorator to add live streaming details.
type LiveStreamingDecorator struct {
	CricketGameDecorator
	platform string
}

// NewLiveStreamingDecorators initializes a new LiveStreamingDecorator for the specified game and platform.
func NewLiveStreamingDecorators(game CricketGame, platform string) *LiveStreamingDecorator {
	return &LiveStreamingDecorator{
		CricketGameDecorator: CricketGameDecorator{game: game}, // Initialize the base decorator.
		platform:             platform,                         // Set the streaming platform.
	}
}

// Play method overrides the Play method to indicate live streaming details.
func (live *LiveStreamingDecorator) Play() string {
	return live.game.Play() + fmt.Sprintf(" streaming live on %s", live.platform)
}

// GetScore method overrides the GetScore method to indicate live watching option.
func (live *LiveStreamingDecorator) GetScore() string {
	return live.game.GetScore() + " - Watch live"
}
