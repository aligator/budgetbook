package intc

import (
	"errors"
)

// In general, intc.Command is used to depict any feature of the application
// and is general purpose, package-wide interchangeable command that also
// serves as a vehicle for parsed data.

// Command represents a certain functionality of the application. It does not
// only provide information about the technical CLI command itself, but also
// defines which data (see Params) is required to execute it appropriately.
type Command struct {
	Use     string
	Help    string
	Params  map[string]*Param
	Options map[string]*Option
	Run     func(self *Command) error
}

// Adds a new param to a given command. Since duplicate param names aren't
// valid, an error will be returned if the param name already exists.
func (cmd *Command) AddParam(p *Param) error {
	if cmd.Params == nil {
		cmd.Params = make(map[string]*Param)
	}
	for key, param := range cmd.Params {
		if key == p.Name || param.Name == p.Name {
			return errors.New("param name already exists")
		}
	}
	cmd.Params[p.Name] = p
	return nil
}

// Adds a new option to a given command. Since duplicate option names aren't
// valid, an error will be returned if the option name already exists.
func (cmd *Command) AddOption(o *Option) error {
	if cmd.Options == nil {
		cmd.Options = make(map[string]*Option)
	}
	for key, opt := range cmd.Options {
		if key == o.Name || opt.Name == o.Name {
			return errors.New("option name already exists")
		}
	}
	cmd.Options[o.Name] = o
	return nil
}
