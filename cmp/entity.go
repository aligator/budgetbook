package cmp

// Entity represents a general purpose model. It is mainly used as an
// anonymous member to carry out the polymorphism mechanism, meaning that
// Entity is representative for all models stored in the cmp package.
type Entity struct {
	ID string
}
