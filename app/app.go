package app

import (
	"budgetBook/cli"
	"budgetBook/intc"
	"budgetBook/persist"
	"fmt"
)

// app represents the application itself. Essentially, it consists of a root
// command, a set of sub-commands, a CLI Proxy and a database implementation.
// The domain specific business logic is depicted in the respective command.
type app struct {
	RootCmd *intc.Command
	CmdSet  []*intc.Command
	DB      persist.Database
	Proxy   cli.Proxy
}

// Creates a new instance of app and initializes all of its components.
func New() *app {
	rootCmd, cmdSet := buildCommandSet()
	a := &app{
		RootCmd: rootCmd,
		CmdSet:  cmdSet,
		DB:      &persist.Bolt{},
		Proxy:   cli.Cobra,
	}
	_ = a.DB.Setup()
	a.Proxy.Setup(a.RootCmd, a.CmdSet)
	return a
}

// Runs the application by parsing the CLI input, transforming it into an
// interchangeable command and executing that specific command.
func (a *app) Run() {
	execCmd := a.Proxy.Parse()
	for _, p := range execCmd.Params {
		if p.Store == "" {
			p.Store = "-"
		}
		fmt.Println(p.Name, ":", p.Store)
	}
	for _, opt := range execCmd.Options {
		store := "false"
		if opt.Store {
			store = "true"
		}
		fmt.Println(opt.Name, ":", store)
	}
}
