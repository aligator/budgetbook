package cli

import (
	"budgetBook/intc"
	"github.com/spf13/cobra"
)

var Cobra = &_cobra{}

// Proxy defines a method set for transforming an intc.Command into a
// concrete CLI command provided by the corresponding library.
type Proxy interface {
	// Initializes the Proxy by transforming a given command set.
	Setup(rootCmd *intc.Command, cs []*intc.Command)
	// Transforms a given intc.Command into a concrete command type.
	transform(cmd *intc.Command) *container
	// Parses the CLI input and returns the data as a intc.Command.
	Parse() *intc.Command
}

// container basically wraps the actual command and holds a simple map
// of pointers returned by the transformed flags.
type container struct {
	Cmd       *cobra.Command
	FlagStore map[string]interface{}
}

// Cobra is only one of several possible implementations of Proxy. It
// refers to the Cobra library (see github.com/spf13/cobra) and therefore
// holds instances of cobra.Command as struct member variables.
type _cobra struct {
	Root   *container
	CtrSet []*container
}

// Implements Proxy.Setup().
func (c *_cobra) Setup(rootCmd *intc.Command, cmds []*intc.Command) {
	c.Root = c.transform(rootCmd)
	for _, cmd := range cmds {
		tfContainer := c.transform(cmd)
		c.CtrSet = append(c.CtrSet, tfContainer)
		c.Root.Cmd.AddCommand(tfContainer.Cmd)
	}
}

// Implements Proxy.transform().
func (c *_cobra) transform(cmd *intc.Command) *container {
	cobraCmd := &cobra.Command{
		Use: cmd.Use,
	}
	fs := make(map[string]interface{})
	cobraCmd.SetHelpTemplate(cmd.Help)
	for _, opt := range cmd.Options {
		fs[opt.Name] = cobraCmd.Flags().StringP(opt.Name, opt.Shorthand, opt.DefVal, opt.Help)
	}
	return &container{
		Cmd:       cobraCmd,
		FlagStore: fs,
	}
}

// Implements Proxy.Parse().
func (c *_cobra) Parse() *intc.Command {
	c.Root.Cmd.Execute()
	return nil
}
