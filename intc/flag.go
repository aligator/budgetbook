package intc

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

// Creates a new instance of Param with a non-empty store. Due to the lack
// of need of any other values like shorthand or default value, only the
// param name is required.
func NewParamByStore(name, store string) *Param {
	p := &Param{
		flag:  &flag{Name: name},
		Store: store,
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

// Creates a new instance of Option with a non-empty store. Due to the lack
// of need of any other values like shorthand or default value, only the
// option name is required.
func NewOptionByStore(name string, store bool) *Option {
	o := &Option{
		flag:  &flag{Name: name},
		Store: store,
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
