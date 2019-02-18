package cmp

// cmp.Type describes if the component is whether associated with an income
// or with an expenditure. Despite being used in transaction.Tx, Type could
// be used in any Entity implementation to depict the financial context.
type Type int

const (
	// Represents a financial income or a context related with it.
	Inc Type = iota
	// Represents a financial expenditure or a context related with it.
	Exp Type = iota
)

// Entity represents any model stored in the cmp package. It is mainly used
// to carry out polymorphism and to prescribe the marshalling functions.
type Entity interface {
	// Forces the entity implementation to provide an ID for persistence.
	ID() string
	// Converts the entity into JSON, which is stored as a byte slice.
	MarshalJSON() ([]byte, error)
	// Fills an empty entity instance with values from JSON.
	UnmarshalJSON(b []byte) error
}
