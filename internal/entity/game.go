package entity

import "time"

// GameStatus represents the possible states of a Bingo game.
type GameStatus string

const (
	StatusWaiting   GameStatus = "WAITING"
	StatusOngoing   GameStatus = "ONGOING"
	StatusCompleted GameStatus = "COMPLETED"
)

// Player represents a user participating in a specific game instance.
// It holds the user's identifier and the specific cards they are playing.
type Player struct {
	UserID             int64 // References entity.User.ID
	CardPaletteNumbers []int // Slice of entity.Card.PaletteNumber assigned to this player for this game
}

// WinnerInfo holds details about the game's winner, if any.
type WinnerInfo struct {
	UserID            int64 // References entity.User.ID
	CardPaletteNumber int   // The specific winning card's palette number
	VerifiedAt        time.Time
}

// Game represents a single instance of a Bingo game.
type Game struct {
	ID               string      // Unique identifier for this game instance
	RoomID           string      // Identifier for the associated Room (e.g., Room.ID or Room.Name)
	AvailableNumbers []int       // Pool of numbers available to be drawn for this game
	DrawnNumbers     []int       // Numbers that have been called/drawn
	StartTime        time.Time   // When the game officially started allowing players or drawing
	EndTime          *time.Time  // Pointer: Time the game concluded (nil if ongoing/waiting)
	Status           GameStatus  // Current status of the game
	Players          []Player    // List of players currently in the game
	Winner           *WinnerInfo // Pointer: Winner details (nil if no winner yet)
	CreatedAt        time.Time   // Timestamp of game record creation
	UpdatedAt        time.Time   // Timestamp of last game record update
}

// Example intrinsic validation method
func (g *Game) HasPlayer(userID int64) bool {
	for _, p := range g.Players {
		if p.UserID == userID {
			return true
		}
	}
	return false
}
