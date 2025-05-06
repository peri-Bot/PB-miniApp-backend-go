package entity

import "time"

// TransactionType defines the type of financial transaction.
type TransactionType string

const (
	TxDeposit    TransactionType = "DEPOSIT"
	TxWithdrawal TransactionType = "WITHDRAWAL"
	TxWin        TransactionType = "WIN"
	TxLoss       TransactionType = "LOSS"
)

// Transaction represents a record of balance change for a user.
type Transaction struct {
	Type      TransactionType
	Amount    int64     // Amount in smallest currency unit (e.g., cents)
	GameID    *string   // Pointer: Associated Game.ID, optional (nil for non-game transactions)
	Timestamp time.Time // When the transaction occurred
}

// User represents a registered player in the system.
type User struct {
	ID               int64  // Primary identifier (from Telegram user_id)
	ChatID           int64  // Telegram chat_id associated with the user
	PhoneNumber      string // User's phone number
	LanguageCode     string // Preferred language code (e.g., "en", "es")
	Username         string // User's Telegram username (or name)
	Balance          int64  // User's balance in smallest currency unit (e.g., cents)
	RegistrationDate time.Time
	Transactions     []Transaction // History of financial transactions
	LastActiveAt     time.Time     // Timestamp of the user's last known activity
	JoinedRoomIDs    []string      // Slice of Room IDs the user is currently associated with (if needed)
	CreatedAt        time.Time     // Timestamp of user record creation
	UpdatedAt        time.Time     // Timestamp of last user record update
}

// Example intrinsic validation method
func (u *User) CanAfford(amount int64) bool {
	return u.Balance >= amount
}
