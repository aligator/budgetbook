package app

import (
	"budgetBook/persistence"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type app struct {
	DB      persistence.Database
	RootCmd *cobra.Command
	Config  *viper.Viper
}

func New() *app {
	bolt := &persistence.Bolt{}
	if err := bolt.Load(); err != nil {
		return nil
	}
	return &app{
		DB:      bolt,
		RootCmd: RootCmd,
	}
}

func (a *app) Run() {
	if !a.RootCmd.HasSubCommands() {
		a.bindCommands(Cmds)
	}
	a.RootCmd.Execute()
}

func (a *app) bindCommands(cmds []*cobra.Command) {
	for i, _ := range cmds {
		a.RootCmd.AddCommand(cmds[i])
	}
}
