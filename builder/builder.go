package builder

/* Builder pattern seperates the construction of a complex object from its representation.
 */
// CricketGame represents a cricket game configuration
type CricketGame struct {
	TeamName    string
	Overs       int
	Location    string
	Players     []string
	Format      string
	Umpires     []string
	HasDrs      bool
	FloodLights bool
}

// CricketGameBuilder is the builder for CricketGame
type CricketGameBuilder struct {
	game *CricketGame
}

// NewCricketGameBuilder intilizes a new builder for cricket came
func NewCricketGameBuilder(teamName string) *CricketGameBuilder {
	return &CricketGameBuilder{
		game: &CricketGame{
			TeamName:    teamName,
			Overs:       20,
			Location:    "Unknown",
			Players:     []string{},
			Format:      "T20",
			Umpires:     []string{},
			HasDrs:      false,
			FloodLights: false,
		},
	}
}

// Here we are going use chaining method where we can use it
// SetOvers will sets the overs and format of the game
func (b *CricketGameBuilder) SetOvers(overs int) *CricketGameBuilder {
	b.game.Overs = overs
	if overs == 50 {
		b.game.Format = "ODI"
	} else if overs == 90 {
		b.game.Format = "Test"
	}
	return b
}

// SetLocation sets the location for the game
func (b *CricketGameBuilder) SetLocation(location string) *CricketGameBuilder {
	b.game.Location = location
	return b
}

// Addplayer adds a player to the game
func (b *CricketGameBuilder) AddPlayers(player string) *CricketGameBuilder {
	b.game.Players = append(b.game.Players, player)
	return b
}

// Addumpire adds an umpire to the game
func (b *CricketGameBuilder) AddUmpires(umpires string) *CricketGameBuilder {
	b.game.Umpires = append(b.game.Umpires, umpires)
	return b
}

// SetDRS addds sets the DRS to the game
func (b *CricketGameBuilder) SetDRS(hasDRS bool) *CricketGameBuilder {
	b.game.HasDrs = hasDRS
	return b
}

// SetFloodLights enables or disabled flood lights for the game
func (b *CricketGameBuilder) SetFloodLights(hasFloodLights bool) *CricketGameBuilder {
	b.game.FloodLights = hasFloodLights
	return b
}

// Build finalizes the game configuration and returns it
func (b *CricketGameBuilder) Build() *CricketGame {
	return b.game
}
