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
		Params:  nil,
		Options: nil,
		Run:     nil,
	}
	addCategory := &intc.Command{
		Use:  "add-cat",
		Help: ``,
		Params: []*intc.Param{
			intc.NewParam("name", "n", ``, ""),
			intc.NewParam("budget", "b", ``, ""),
		},
		Options: []*intc.Option{
			intc.NewOption("is-inc", "i", ``, false),
			intc.NewOption("is-capped", "c", ``, false),
		},
	}
	cmdSet := []*intc.Command{addCategory}
	return rootCmd, cmdSet
}
