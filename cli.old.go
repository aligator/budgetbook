package cli

import (
	"budgetBookFlex/app"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type C struct {
	Actual interface{}
}

type UI struct {
	RootCmd      *C
	Cmds         []*C
	TfHandler    func(cmd *app.Command) interface{}
	ParseHandler func(rootCmd *C, cmds []*C) *app.Command
}

func (m *UI) Transform(rootCmd *app.Command, cmds []*app.Command) {
	tfCommand := m.TfHandler(rootCmd)
	m.RootCmd = &C{
		Actual: tfCommand,
	}
	for _, appCmd := range cmds {
		tfAppCmd := m.TfHandler(appCmd)
		transformed := &C{
			Actual: tfAppCmd,
		}
		m.Cmds = append(m.Cmds, transformed)
	}
}

func (m *UI) Parse() *app.Command {
	return m.ParseHandler(m.RootCmd, m.Cmds)
}

// noinspection GoUnresolvedReference
var Cobra = &UI{
	RootCmd: &C{Actual: &cobra.Command{}},
	Cmds:    []*C{},
	TfHandler: func(cmd *app.Command) interface{} {
		cobraCmd := &cobra.Command{
			Use: cmd.Use,
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("So weit so gut.")
			},
		}
		for _, f := range cmd.Flags {
			cobraCmd.Flags().AddFlag(&pflag.Flag{
				Name:      f.Name,
				Shorthand: f.Short,
				DefValue:  f.StdVal,
			})
		}
		return cobraCmd
	},
	ParseHandler: func(rootCmd *C, cmds []*C) *app.Command {
		_, isCobra := rootCmd.Actual.(cobra.Command)
		if !isCobra {
			return nil
		}
		for _, cmd := range cmds {
			if _, ok := cmd.Actual.(cobra.Command); ok {
				rootCmd.Actual.AddComand(cmd)
			}
		}
		rootCmd.Execute()
		return nil
	},
}
