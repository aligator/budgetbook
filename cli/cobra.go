package cli

import (
	"budgetBook/intc"
	"github.com/spf13/cobra"
)

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
		// it to the container set.
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
	fs := make(map[string]*string)
	for _, opt := range cmd.Params {
		// Map the returned pointer to the flag value against its name.
		// Add the flag shorthand only if the shorthand is not empty.
		if opt.Shorthand != "" {
			fs[opt.Name] = cobraCmd.Flags().StringP(opt.Name, opt.Shorthand, opt.DefVal, opt.Help)
		} else {
			fs[opt.Name] = cobraCmd.Flags().String(opt.Name, opt.DefVal, opt.Help)
		}
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
			cmd.AddFlag(&intc.Param{
				Name:  key,
				Store: *val,
			})
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
		return c.inverse(c.findInCtrSet(cmd))
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
