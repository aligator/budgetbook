package app

import (
	"budgetBook/cli"
	"budgetBook/intc"
	"budgetBook/persist"
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

// Runs the application by parsing the CLI input, transforming it into a
// interchangeable command and executing that specific command.
func (a *app) Run() {
	_ = a.Proxy.Parse()
}

// Builds the entire command set including the root command and returns
// interchangeable, general purpose commands that can be transformed by
// an cli.Proxy for parsing the CLI input.
func buildCommandSet() (*intc.Command, []*intc.Command) {
	rootCmd := &intc.Command{
		Use:     "budgetbook",
		Help:    ``,
		Options: nil,
		Run:     nil,
	}
	addCategory := &intc.Command{
		Use:  "add-cat",
		Help: ``,
		Options: []*intc.Flag{
			{
				Name:      "name",
				Shorthand: "n",
				DefVal:    "",
			},
		},
	}
	cmdSet := []*intc.Command{addCategory}
	return rootCmd, cmdSet
}
