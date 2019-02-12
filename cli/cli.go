package cli

import (
	"budgetBook/intc"
)

var Cobra = &_cobra{}

// Explaining the BudgetBook CLI mechanism:
// ...
// ...
// ...

// Proxy defines a method set for transforming an intc.Command into a
// concrete CLI command provided by the corresponding library.
type Proxy interface {
	// Initializes the Proxy by transforming a given command set.
	Setup(rootCmd *intc.Command, cs []*intc.Command)
	// Transforms a given intc.Command into a concrete command type.
	transform(cmd *intc.Command) *container
	// Describes the inverse function of transform().
	inverse(ctn *container) *intc.Command
	// Parses the CLI input and returns the data as a intc.Command.
	Parse() *intc.Command
}

// container basically wraps the actual command and holds a simple map of
// pointers returned by the transformed flags. After executing the command,
// the store's values may be used for the intc.Command returned by Parse().
type container struct {
	Cmd       interface{}
	FlagStore map[string]*string
}
