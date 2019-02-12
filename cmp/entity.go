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

// The Repository interface makes sure that every repository for the cmp
// types implement the essential methods listed below.
type Repository interface {
	// Finds an entity by id.
	Find(id string) *Entity
	// Returns all entities of the table the repository is responsible for.
	FindAll() []*Entity
	// Inserts a given entity into the corresponding table. If the entity
	// already exists, an error will be returned.
	Insert(e *Entity) error
	// Updates a given entity. If it doesn't exist yet, it will be inserted.
	Update(e *Entity) error
	// Deletes a given entity from its containing table.
	Delete(e *Entity) error
}
