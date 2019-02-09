package cmp

// cmp.Type describes if the component is whether associated with an income
// or with an expenditure.
type Type int

const (
	// Represents an financial income.
	Inc Type = iota
	// Represents an financial expenditure.
	Exp Type = iota
)

// Entity represents a general purpose model. It is mainly used as an
// anonymous member to carry out the polymorphism mechanism, meaning that
// Entity is representative for all models stored in the cmp package.
type Entity struct {
	ID string `json:"id"`
}
