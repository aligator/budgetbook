package intc

import (
	"budgetBookArch/conf"
	"errors"
)

// In general, intc.Command is used to depict any feature of the application
// and is general purpose, package-wide interchangeable command that also
// serves as a vehicle for parsed data.
//
// Command represents a certain functionality of the application. It does not
// only provide information about the technical CLI command itself, but also
// defines which data (see Params) is required to handle it appropriately.
type Command struct {
	Use     string
	Help    string
	Params  map[string]*Param
	Options map[string]*Option
	Run     func(cmd *Command) error
}

// Adds a new param to a given command. Since duplicate param names aren't
// valid, an error will be returned if the param name already exists.
func (cmd *Command) AddParam(p *Param) error {
	if cmd.Params == nil {
		cmd.Params = make(map[string]*Param)
	}
	for _, param := range cmd.Params {
		if param.Name == p.Name {
			return errors.New(conf.CmdParamAlreadyExists)
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
	for _, opt := range cmd.Options {
		if opt.Name == o.Name {
			return errors.New(conf.CmdOptionAlreadyExists)
		}
	}
	cmd.Options[o.Name] = o
	return nil
}

// Returns the value of a parameter with the specified key by de-referencing
// the stored pointer. In case the key does not exist in the map of parameters,
// the zero value will be returned.
func (cmd *Command) P(key string) string {
	if p, ok := cmd.Params[key]; ok {
		return *p.Val
	}
	return ""
}

// Returns the value of a option with the specified key by de-referencing
// the stored pointer. In case the key does not exist in the map of options,
// the zero value will be returned.
func (cmd *Command) Opt(key string) bool {
	if opt, ok := cmd.Options[key]; ok {
		return *opt.Val
	}
	return false
}
