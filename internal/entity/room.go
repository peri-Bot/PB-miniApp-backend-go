package entity

import "time"

// Room represents a virtual space where games are organized and played.
type Room struct {
	ID               string    // Unique identifier for the room (likely the Name)
	Name             string    // Display name of the room
	StakeAmount      int64     // Required amount (smallest unit) to join a game in this room
	MaxPlayers       int       // Maximum number of players allowed in a game within this room
	CurrentPlayerIDs []int64   // Slice of User IDs currently present or associated with the room
	CreatedAt        time.Time // Timestamp of room record creation
	UpdatedAt        time.Time // Timestamp of last room record update
}

// Example intrinsic validation method
func (r *Room) IsFull() bool {
	return len(r.CurrentPlayerIDs) >= r.MaxPlayers
}
