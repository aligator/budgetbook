package app

import (
	"budgetBookArch/cli"
	"budgetBookArch/intc"
	"budgetBookArch/persist"
)

type app struct {
	Root     *intc.Command
	CmdSet   []*intc.Command
	Mediator cli.CLI
	DB       persist.Database
}

func (a *app) Run() {
	//defer a.DB.Close()
	exec, _ := a.Mediator.Parse()

	if exec != nil && exec.Run != nil {
		exec.Run(exec)
	}
}

func New() *app {
	root, cmdSet := buildCmds()
	a := &app{
		Root:     root,
		CmdSet:   cmdSet,
		Mediator: cli.New(),
		DB:       persist.New(),
	}
	a.Mediator.Register(a.Root, a.CmdSet)
	return a
}
