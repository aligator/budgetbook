package app

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: `budgetbook <Command>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`Welcome to BudgetBook`)
	},
}

var Cmds = []*cobra.Command{
	&cobra.Command{},
	&cobra.Command{},
}
