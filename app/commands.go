package app

import (
	"budgetBook/cmp/category"
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
		Params: map[string]*intc.Param{
			"name":   intc.NewParam("name", "n", ``, ""),
			"budget": intc.NewParam("budget", "b", ``, "0"),
		},
		Options: map[string]*intc.Option{
			"is-inc":    intc.NewOption("is-inc", "i", ``, false),
			"is-capped": intc.NewOption("is-capped", "c", ``, false),
		},
		Run: func(self *intc.Command) error {
			return category.Create(self)
		},
	}
	displayCategories := &intc.Command{
		Use: "show-cats",
		Help: ``,
		Params: nil,
		Options: map[string]*intc.Option{
			"inc-only": intc.NewOption("inc-only", "i", ``, false),
		},
		Run: func(self *intc.Command) error {
			return category.ShowAll()
		},
	}
	cmdSet := []*intc.Command{addCategory, displayCategories}
	return rootCmd, cmdSet
}
