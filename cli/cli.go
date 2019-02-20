package cli

import "budgetBook/intc"

// CLI defines a method set for transforming an general-purpose intc.Command
// into a concrete command type of any CLI library. These concrete commands
// are wrapped into containers (see below).
// In short, a CLI implementation's task is to parse the user input and store
// the gathered information in the associated intc.Command from the container.
type CLI interface {
	// Initializes the CLI library by transforming the given command set into
	// specific command types that are used by the library.
	Register(root *intc.Command, cmdSet []*intc.Command)
	// Parses the user input and returns the data as an interchangeable command.
	Parse() (*intc.Command, error)
	// Transforms an interchangeable command into a concrete command type.
	transform(cmd *intc.Command) *container
}

// container wraps a concrete command type of the used CLI library next to
// a intc.Command the concrete command is associated with. The association
// results from the transformation that was performed in transform().
type container struct {
	Cmd      interface{}
	AssocCmd *intc.Command
}

// Creates a new instance of a CLI implementation and returns a pointer to
// that instance. Any implementation should be an unexported type to force
// the use of this factory.
func New() CLI {
	c := &_cobra{}
	return c
}
