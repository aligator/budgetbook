package app

import (
	"budgetBook/intc"
)

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
		Options: []*intc.Param{
			{Name: "name", Shorthand: "n", Help: ``, DefVal: ""},
			{Name: "inc", Shorthand: "", Help: ``, DefVal: ""},
			{Name: "capped", Shorthand: "c", Help: ``, DefVal: ""},
			{Name: "budget", Shorthand: "b", Help: ``, DefVal: ""},
		},
	}
	cmdSet := []*intc.Command{addCategory}
	return rootCmd, cmdSet
}
