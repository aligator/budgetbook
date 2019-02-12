package intc

import (
	"errors"
)

// Command represents a certain functionality of the application. It does not
// only provide information about the technical CLI command itself, but also
// defines which data (see Options) is required to execute it appropriately.
type Command struct {
	Use     string
	Help    string
	Options []*Param
	Run     func(options []*Param) error
}

// Param represents one of the command's options. Holding a shorthand name
// next to a default value makes Param suitable for using it as a CLI
// command flag. Since its value is whether a string or a bool, it must
// be casted or parsed to any different type if necessary.
type Param struct {
	Name      string
	Shorthand string
	IsBool    bool
	Help      string
	DefVal    string
	Store     interface{}
}

// Adds a new flag to a given command. Since duplicate flag names aren't
// valid, an error will be returned if the flag name already exists.
func (cmd *Command) AddFlag(f *Param) error {
	for _, opt := range cmd.Options {
		if opt.Name == f.Name {
			return errors.New("flag name already exists")
		}
	}
	cmd.Options = append(cmd.Options, f)
	return nil
}
