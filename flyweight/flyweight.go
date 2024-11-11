package flyweight

// PlayerType represents the intrisic state of a cricket player
type PlayerType struct {
	role       string
	skillLevel string
	baseStats  map[string]int
}

// Player represents the context combining intrinsic and extrinsic states
type Player struct {
	playerType  *PlayerType
	name        string
	position    string
	currentForm int
}

// PlayerFactory creates a new player factory
type PlayerFactory struct {
	playerTypes map[string]*PlayerType
}

// NewPlayerFactory creates a new player factory
func NewPlayerFactory() *PlayerFactory {
	return &PlayerFactory{
		playerTypes: make(map[string]*PlayerType),
	}
}

// GetPlayerType retrieves or creates a PlayerType flyweight
func (f *PlayerFactory) GetPlayerType(role string, skillLevel string) *PlayerType {
	key := role + "-" + skillLevel

	if pt, exists := f.playerTypes[key]; exists {
		return pt
	}

	newPlayerType := &PlayerType{
		role:       role,
		skillLevel: skillLevel,
		baseStats:  getBaseStats(role, skillLevel),
	}
	f.playerTypes[key] = newPlayerType
	return newPlayerType
}

func NewPlayer(playerType *PlayerType, name string, position string) *Player {
	return &Player{
		playerType:  playerType,
		name:        name,
		position:    position,
		currentForm: 100,
	}
}

// Helper function to get base status
func getBaseStats(role string, skillLevel string) map[string]int {
	stats := make(map[string]int)
	switch role {
	case "Batsman":
		stats["batting"] = 80
		stats["bowling"] = 20
		stats["fielding"] = 60
	case "Bowler":
		stats["batting"] = 30
		stats["bowling"] = 85
		stats["fielding"] = 50
	case "AllRounder":
		stats["batting"] = 65
		stats["bowling"] = 65
		stats["fielding"] = 70
	}
	multiplier := 1.0
	switch skillLevel {
	case "Elite":
		multiplier = 1.2
	case "Regular":
		multiplier = 1.0
	case "Rookie":
		multiplier = 0.8
	}

	for stat := range stats {
		stats[stat] = int(float64(stats[stat]) * multiplier)
	}
	return stats
}
