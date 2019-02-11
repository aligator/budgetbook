package cmp

// cmp.Type describes if the component is whether associated with an income
// or with an expenditure.
type Type int

const (
	// Represents a financial income.
	Inc Type = iota
	// Represents a financial expenditure.
	Exp Type = iota
)

// Entity represents any model stored in the cmp package. It is mainly used
// to carry out the polymorphism mechanism and to prescribe the marshalling
// methods.
type Entity interface {
	// Forces the entity implementation to provide an ID.
	ID() string
	// Converts the entity into JSON.
	MarshalJSON() ([]byte, error)
	// Fills an empty entity with values from JSON.
	UnmarshalJSON(b []byte) error
}
