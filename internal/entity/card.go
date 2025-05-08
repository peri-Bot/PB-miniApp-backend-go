package entity

// CardNumbers represents the B-I-N-G-O columns on a card.
// The 'N' column traditionally has a free space, which might be represented
// by a convention (e.g., 0) or handled in generation/validation logic.
type CardNumbers struct {
	B []int
	I []int
	N []int // Expect 4 numbers + potentially a 'free space' marker
	G []int
	O []int
}

// Card represents a single Bingo card identified by its palette number.
type Card struct {
	PaletteNumber int // Unique identifier for a card template
	Numbers       CardNumbers
}

// Potential intrinsic validation (example - can be expanded)
func (cn *CardNumbers) IsValidStructure() bool {
	// Basic check - real validation might be more complex (e.g., number ranges)
	return len(cn.B) == 5 && len(cn.I) == 5 && len(cn.N) == 5 && len(cn.G) == 5 && len(cn.O) == 5
}
