package intc

import (
	"errors"
)

// Command represents a certain functionality of the application. It does not
// only provide information about the technical CLI command itself, but also
// defines which data (see Params) is required to execute it appropriately.
type Command struct {
	Use    string
	Help   string
	Params []*Param
	Run    func(options []*Param) error
}

// Param represents one of the command's parameters. Holding a shorthand name
// next to a default value makes Param suitable for using it as a CLI
// command flag.
type Param struct {
	Name      string
	Shorthand string
	Help      string
	DefVal    string
	Store     interface{}
}

// Adds a new flag to a given command. Since duplicate flag names aren't
// valid, an error will be returned if the flag name already exists.
func (cmd *Command) AddFlag(f *Param) error {
	for _, opt := range cmd.Params {
		if opt.Name == f.Name {
			return errors.New("flag name already exists")
		}
	}
	cmd.Params = append(cmd.Params, f)
	return nil
}
