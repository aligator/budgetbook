package intc

type flag struct {
	Name      string
	Shorthand string
	Help      string
}

type Param struct {
	*flag
	DefaultVal string
	Val        *string
}

type Option struct {
	*flag
	DefaultVal bool
	Val        *bool
}

func newFlag(name, shorthand, help string) *flag {
	f := &flag{
		Name:      name,
		Shorthand: shorthand,
		Help:      help,
	}
	return f
}

func NewParam(name, shorthand, help, defaultVal string, val *string) *Param {
	p := &Param{
		flag:       newFlag(name, shorthand, help),
		DefaultVal: defaultVal,
		Val:        val,
	}
	return p
}

func NewOption(name, shorthand, help string, defaultVal bool, val *bool) *Option {
	o := &Option{
		flag:       newFlag(name, shorthand, help),
		DefaultVal: defaultVal,
		Val:        val,
	}
	return o
}
