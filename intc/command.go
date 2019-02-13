package intc

import (
	"errors"
)

// Command represents a certain functionality of the application. It does not
// only provide information about the technical CLI command itself, but also
// defines which data (see Params) is required to execute it appropriately.
type Command struct {
	Use     string
	Help    string
	Params  []*Param
	Options []*Option
	Run     func(options []*Param) error
}

// flag is a struct component for command parameters and options. Holding a
// shorthand name next to a default value makes embedding types suitable for
// using them as a CLI command flag.
type flag struct {
	Name      string
	Shorthand string
	Help      string
}

// Param represents one of the command's parameters that holds data for further
// processing.
type Param struct {
	*flag
	DefVal string
	Store  string
}

// Option represents a flag of type bool, such as --verbose or --all.
type Option struct {
	*flag
	DefVal bool
	Store  bool
}

// Creates a new instance of Param and returns a pointer to that instance.
// Despite Param is an exported type, it is recommended to use this factory
// since it includes the initialization of an unexported component.
func NewParam(name, shorthand, help, defVal string) *Param {
	p := &Param{
		flag:   newFlag(name, shorthand, help),
		DefVal: defVal,
	}
	return p
}

// Creates a new instance of Option and returns a pointer to that instance.
// Despite Param is an exported type, it is recommended to use this factory
// since it includes the initialization of an unexported component.
func NewOption(name, shorthand, help string, defVal bool) *Option {
	o := &Option{
		flag:   newFlag(name, shorthand, help),
		DefVal: defVal,
	}
	return o
}

// Creates a new instance of flag and returns a pointer to that instance.
func newFlag(name, shorthand, help string) *flag {
	f := &flag{
		Name:      name,
		Shorthand: shorthand,
		Help:      help,
	}
	return f
}

// Adds a new param to a given command. Since duplicate param names aren't
// valid, an error will be returned if the param name already exists.
func (cmd *Command) AddParam(p *Param) error {
	for _, param := range cmd.Params {
		if param.Name == p.Name {
			return errors.New("param name already exists")
		}
	}
	cmd.Params = append(cmd.Params, p)
	return nil
}

// Adds a new option to a given command. Since duplicate option names aren't
// valid, an error will be returned if the option name already exists.
func (cmd *Command) AddOption(o *Option) error {
	for _, opt := range cmd.Options {
		if opt.Name == o.Name {
			return errors.New("option name already exists")
		}
	}
	cmd.Options = append(cmd.Options, o)
	return nil
}
