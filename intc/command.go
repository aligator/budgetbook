package intc

import "errors"

// Command represents a certain functionality of the application. It does not
// only provide information about the technical CLI command itself, but also
// defines which data (see Options) is required to execute it appropriately.
type Command struct {
	Use     string
	Help    string
	Options []*Flag
	Run     func() error
}

// Flag represents one of the command's options. Holding a shorthand name
// next to a default value makes Flag suitable for using it as a CLI
// command flag.
type Flag struct {
	Name      string
	Shorthand string
	DefVal    string
	Store     interface{}
}

// Adds a new flag to a given command. Since duplicate flag names aren't
// valid, an error will be returned if the flag name already exists.
func (cmd *Command) AddFlag(f *Flag) error {
	for _, opt := range cmd.Options {
		if opt.Name == f.Name {
			return errors.New("flag name already exists")
		}
	}
	cmd.Options = append(cmd.Options, f)
	return nil
}
