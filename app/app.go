package app

import (
	"budgetBook/cli"
	"budgetBook/intc"
	"budgetBook/persist"
	"fmt"
)

// app represents the application itself. Essentially, it consists of a root
// command next to a set of sub-commands, a CLI and a database implementation.
// While it initiates the data flow, the domain specific business logic is
// depicted in the respective command.
type app struct {
	Root     *intc.Command
	CmdSet   []*intc.Command
	Mediator cli.CLI
	DB       persist.Database
}

// Runs the application by parsing the CLI input and executing the gathered
// command which delegates the processing to the provided handler.
func (a *app) Run() {
	//defer a.DB.Close()
	exec, _ := a.Mediator.Parse()

	// Check if a handler has been assigned to the executed command.
	if exec != nil && exec.Run != nil {
		if err := exec.Run(exec); err != nil {
			fmt.Println(err)
		}
	}
}

// Creates a new instance of app and returns a pointer to that instance.
// This factory function is required since the zero value state of app is
// not usable without registering the commands in the CLI mediator.
func New() *app {
	root, cmdSet := buildCmds()
	a := &app{
		Root:     root,
		CmdSet:   cmdSet,
		Mediator: cli.New(),
		DB:       nil,
	}
	a.Mediator.Register(a.Root, a.CmdSet)
	return a
}
