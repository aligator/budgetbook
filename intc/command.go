package intc

import (
	"budgetBookArch/str"
	"errors"
)

type Command struct {
	Use     string
	Help    string
	Params  map[string]*Param
	Options map[string]*Option
	Run     func(cmd *Command) error
}

func (cmd *Command) AddParam(p *Param) error {
	if cmd.Params == nil {
		cmd.Params = make(map[string]*Param)
	}
	for _, param := range cmd.Params {
		if param.Name == p.Name {
			return errors.New(str.CmdParamAlreadyExists)
		}
	}
	cmd.Params[p.Name] = p
	return nil
}

func (cmd *Command) AddOption(o *Option) error {
	if cmd.Options == nil {
		cmd.Options = make(map[string]*Option)
	}
	for _, opt := range cmd.Options {
		if opt.Name == o.Name {
			return errors.New(str.CmdOptionAlreadyExists)
		}
	}
	cmd.Options[o.Name] = o
	return nil
}

func (cmd *Command) P(key string) string {
	if p, ok := cmd.Params[key]; ok {
		return *p.Val
	}
	return ""
}

func (cmd *Command) Opt(key string) bool {
	if opt, ok := cmd.Options[key]; ok {
		return *opt.Val
	}
	return false
}
