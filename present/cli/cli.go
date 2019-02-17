package cli

import "budgetBookArch/app/intc"

type CLI interface {
	Register(root *intc.Command, cmdSet []*intc.Command)
	Parse() (*intc.Command, error)
	transform(cmd *intc.Command) *container
}

type container struct {
	Cmd      interface{}
	AssocCmd *intc.Command
}

func New() CLI {
	c := &_cobra{}
	return c
}
