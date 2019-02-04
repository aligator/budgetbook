package cli

import (
	"budgetBook/intc"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var Cobra = &_cobra{}

// Proxy defines a method set for transforming an intc.Command into a
// concrete CLI command provided by the corresponding library.
type Proxy interface {
	Setup(rootCmd *intc.Command, cs []*intc.Command) // Initializes the Proxy by transforming a given command set.
	transform(cmd *intc.Command) interface{}         // Transforms a given intc.Command into a concrete command type.
	Parse() *intc.Command                            // Parses the CLI input and returns the data as a intc.Command.
}

// Cobra is only one of several possible implementations of Proxy. It
// refers to the Cobra library (see github.com/spf13/cobra) and therefore
// holds instances of cobra.Command as struct member variables.
type _cobra struct {
	RootCmd *cobra.Command
	CmdSet  []*cobra.Command
}

// Implements Proxy.Setup().
func (c *_cobra) Setup(rootCmd *intc.Command, cs []*intc.Command) {
	c.RootCmd, _ = c.transform(rootCmd).(*cobra.Command)
	for _, cmd := range cs {
		tfCmd, _ := c.transform(cmd).(*cobra.Command)
		c.CmdSet = append(c.CmdSet, tfCmd)
		c.RootCmd.AddCommand(tfCmd)
	}
}

// Implements Proxy.transform().
func (c *_cobra) transform(cmd *intc.Command) interface{} {
	cobraCmd := &cobra.Command{
		Use: cmd.Use,
	}
	cobraCmd.SetHelpTemplate(cmd.Help)
	for _, opt := range cmd.Options {
		cobraCmd.Flags().AddFlag(&pflag.Flag{
			Name:      opt.Name,
			Shorthand: opt.Shorthand,
			DefValue:  opt.DefVal,
		})
	}
	return cobraCmd
}

// Implements Proxy.Parse().
func (c *_cobra) Parse() *intc.Command {
	c.RootCmd.Execute()
	return nil
}
