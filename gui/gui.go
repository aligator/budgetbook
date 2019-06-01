package gui

import "budgetBook/intc"

type Gui interface {
	Register(root *intc.Command, cmdSet []*intc.Command)
}

type container struct {
	Gui      interface{}
	AssocCmd *intc.Command
}

func NewWeb() Gui {
	c := &web{}
	return c
}
