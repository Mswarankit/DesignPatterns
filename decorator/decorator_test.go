package decorator

import (
	"strings"
	"testing"
)

func TestBasicCricketGame(t *testing.T) {
	game := NewBasicCricketGame()
	result := game.Play()
	if result != "Playing Cricket" {
		t.Errorf("Expected 'Plaing Cricket', got '%s'", result)
	}
}

func TestCommentratorDecorator(t *testing.T) {
	game := NewBasicCricketGame()
	game.score = 200
	gameWithCommentrator := NewCommentratorDecorator(game, "Ravi Shastri")
	result := gameWithCommentrator.Play()
	if !strings.Contains(result, "with commentary by Ravi Shastri") {
		t.Errorf("Expected Commentry by Ravi Shastri, got '%s'", result)
	}
	score := gameWithCommentrator.GetScore()
	if !strings.Contains(score, "What a fantastic game!") {
		t.Errorf("Expected Cheers from commentry, got '%s'", score)
	}
}

func TestStatisticsDecorator(t *testing.T) {
	game := NewBasicCricketGame()
	gameWithStats := NewStatisticsDecorators(game)

	result := gameWithStats.Play()
	if !strings.Contains(result, "with live statistics") {
		t.Errorf("Expected statistics mention, got '%s'", result)
	}

	score := gameWithStats.GetScore()
	if !strings.Contains(score, "Run Rate:") {
		t.Errorf("Expected run rate in statistics, got '%s'", score)
	}
}

func TestMultipleDecorators(t *testing.T) {
	game := NewBasicCricketGame()

	// Add multiple decorators
	gameWithCommentary := NewCommentratorDecorator(game, "Tony Greig")
	gameWithCommentaryAndStats := NewStatisticsDecorators(gameWithCommentary)
	gameWithAll := NewLiveStreamingDecorators(gameWithCommentaryAndStats, "YouTube")

	result := gameWithAll.Play()

	expectedFeatures := []string{
		"with commentary",
		"with live statistics",
		"streaming live on YouTube",
	}

	for _, feature := range expectedFeatures {
		if !strings.Contains(result, feature) {
			t.Errorf("Expected '%s' in result, got '%s'", feature, result)
		}
	}
}

func TestDecoratorIndependence(t *testing.T) {
	game := NewBasicCricketGame()

	gameWithCommentary := NewCommentratorDecorator(game, "Ravi Shastri")
	gameWithStats := NewStatisticsDecorators(game)

	commentaryResult := gameWithCommentary.Play()
	statsResult := gameWithStats.Play()

	if strings.Contains(commentaryResult, "statistics") {
		t.Error("Commentary decorator should not include statistics")
	}

	if strings.Contains(statsResult, "commentary") {
		t.Error("Statistics decorator should not include commentary")
	}
}
