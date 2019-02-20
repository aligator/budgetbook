package app

import (
	"budgetBook/handle"
	"budgetBook/intc"
)

// Builds the entire command set including the root command. It returns general
// purpose, interchangeable commands that can be transformed by a cli.CLI in
// order to parse the user input.
func buildCmds() (*intc.Command, []*intc.Command) {
	// The root command doesn't perform any action and therefore doesn't take
	// any parameters or options (except for possible help options provided
	// by a cli.Mediator implementation).
	root := &intc.Command{
		Use:  "budgetbook",
		Help: ``,
	}
	// Creates a new transaction category and stores it in the database.
	addCategory := &intc.Command{
		Use:  "add-cat",
		Help: ``,
		Params: map[string]*intc.Param{
			"name":   intc.NewParam("name", "n", ``, "", nil),
			"budget": intc.NewParam("budget", "b", ``, "", nil),
		},
		Options: map[string]*intc.Option{
			"is-inc":    intc.NewOption("is-inc", "i", ``, false, nil),
			"is-capped": intc.NewOption("is-capped", "c", ``, false, nil),
		},
		Run: handle.NewCatController().Create,
	}
	// Prints all stored categories.
	showCategories := &intc.Command{
		Use:    "show-cats",
		Help:   ``,
		Params: nil,
		Options: map[string]*intc.Option{
			"inc-only": intc.NewOption("in-only", "i", ``, false, nil),
		},
		Run: handle.NewCatController().Show,
	}
	// Initialize the actual command set. Aforementioned commands not added to
	// this slice are not returned and hence not registered in the application.
	cmdSet := []*intc.Command{addCategory, showCategories}
	return root, cmdSet
}
