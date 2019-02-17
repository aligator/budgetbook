package app

import (
	"budgetBookArch/app/intc"
)

func buildCmds() (*intc.Command, []*intc.Command) {
	root := &intc.Command{
		Use:  "budgetbook",
		Help: ``,
	}
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
		//Run: handle.NewCatController().Create,
	}
	cmdSet := []*intc.Command{addCategory}
	return root, cmdSet
}
