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
	// Describes the inverse function of transform().
	inverse(ctn *container) *intc.Command
	// Parses the CLI input and returns the data as a intc.Command.
	Parse() *intc.Command
}

// container basically wraps the actual command and holds a simple map of
// pointers returned by the transformed flags. After executing the command,
// the store's values may be used for the intc.Command returned by Parse().
type container struct {
	Cmd       interface{}
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
		// Transform the interchangeable command into a container and add
		// it to the container set
		tfCtn := c.transform(cmd)
		c.CtrSet = append(c.CtrSet, tfCtn)
		// Cast both the root command and the container's command to cobra
		// commands so that AddCommand() can work type safe.
		cobraRootCmd, ok := c.Root.Cmd.(*cobra.Command)
		ctnCmd, ok := tfCtn.Cmd.(*cobra.Command)
		if ok {
			cobraRootCmd.AddCommand(ctnCmd)
		}
	}
}

// Implements Proxy.transform().
func (c *_cobra) transform(cmd *intc.Command) *container {
	cobraCmd := &cobra.Command{
		Use: cmd.Use,
	}
	cobraCmd.SetHelpTemplate(cmd.Help)
	fs := make(map[string]interface{})
	for _, opt := range cmd.Options {
		// Map the returned pointer to the flag value against its name.
		fs[opt.Name] = cobraCmd.Flags().StringP(opt.Name, opt.Shorthand, opt.DefVal, opt.Help)
	}
	return &container{
		Cmd:       cobraCmd,
		FlagStore: fs,
	}
}

// Implements Proxy.inverse().
func (c *_cobra) inverse(ctn *container) *intc.Command {
	if ctnCmd, ok := ctn.Cmd.(*cobra.Command); ok {
		cmd := &intc.Command{
			Use: ctnCmd.Use,
		}
		for key, val := range ctn.FlagStore {
			flag := &intc.Flag{
				Name: key,
				Store: val,
			}
			cmd.AddFlag(flag)
		}
		return cmd
	}
	return nil
}

// Implements Proxy.Parse().
func (c *_cobra) Parse() *intc.Command {
	rootCmd, ok := c.Root.Cmd.(*cobra.Command)
	if ok {
		// Receive the executed command, find and inverse transform it to
		// an interchangeable command filled with data by the flag store.
		cmd, _ := rootCmd.ExecuteC()
		_ = c.findInCtrSet(cmd)
	}
	return nil
}

// findInCtrSet searches a given cobra Command in the container set of the
// parser. Its primary purpose is to determine the specific command that
// has been executed in Proxy.Parse() to access its flag store.
func (c *_cobra) findInCtrSet(cmd *cobra.Command) *container {
	if rootCmd, ok := c.Root.Cmd.(*cobra.Command); ok {
		if rootCmd.Use == cmd.Use {
			return c.Root
		}
	}
	for _, ctr := range c.CtrSet {
		if ctrCmd, ok := ctr.Cmd.(*cobra.Command); ok {
			if ctrCmd.Use == cmd.Use {
				return ctr
			}
		}
	}
	return nil
}
