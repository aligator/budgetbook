package cli

import (
	"budgetBookArch/conf"
	"budgetBookArch/intc"
	"errors"
	"github.com/spf13/cobra"
)

// _cobra is only one of several possible implementations of Mediator. It
// refers to the cobra library (see github.com/spf13/cobra) and therefore
// holds instances of cobra.Command as member variables.
// The preceding underscore is used to avoid package name collisions.
type _cobra struct {
	RootCtr *container
	CtrSet  []*container
}

// Implements CLI.Register().
func (c *_cobra) Register(root *intc.Command, cmdSet []*intc.Command) {
	c.RootCtr = c.transform(root)
	for _, cmd := range cmdSet {
		transformed := c.transform(cmd)
		c.CtrSet = append(c.CtrSet, transformed)
		cobraRoot, ok := c.RootCtr.Cmd.(*cobra.Command)
		ctrCmd, ok := transformed.Cmd.(*cobra.Command)
		if ok {
			cobraRoot.AddCommand(ctrCmd)
		}
	}
}

// Implements CLI.Parse().
func (c *_cobra) Parse() (*intc.Command, error) {
	cobraRoot, ok := c.RootCtr.Cmd.(*cobra.Command)
	if ok {
		executed, err := cobraRoot.ExecuteC()
		if err != nil {
			return nil, err
		}
		ctr, err := c.findInContainers(executed)
		if err != nil {
			return nil, err
		}
		return ctr.AssocCmd, nil

	}
	return nil, errors.New(conf.CmdTypeAssertionFailed)
}

// Implements CLI.transform().
func (c *_cobra) transform(cmd *intc.Command) *container {
	cobraCmd := &cobra.Command{
		Use: cmd.Use,
	}
	for key, p := range cmd.Params {
		paramPtr := cobraCmd.Flags().StringP(p.Name, p.Shorthand, p.DefaultVal, p.Help)
		cmd.Params[key].Val = paramPtr
	}
	for key, opt := range cmd.Options {
		optionPtr := cobraCmd.Flags().BoolP(opt.Name, opt.Shorthand, opt.DefaultVal, opt.Help)
		cmd.Options[key].Val = optionPtr
	}
	return &container{
		Cmd:      cobraCmd,
		AssocCmd: cmd,
	}
}

// Cobra looks for a sub-command that matches the user input and returns that
// command. All cobra commands were initialized in transform() and then stored
// in a container - this function finds that container.
func (c *_cobra) findInContainers(cmd *cobra.Command) (*container, error) {
	// Check if the executed command was the root command. The type assertion
	// is necessary to check the concrete command's use value.
	if rootCmd, ok := c.RootCtr.Cmd.(*cobra.Command); ok {
		if rootCmd.Use == cmd.Use {
			return c.RootCtr, nil
		}
	} else {
		return nil, errors.New(conf.CmdTypeAssertionFailed)
	}
	// Otherwise, all containers are scanned for a matching cobra command.
	for _, ctr := range c.CtrSet {
		if ctrCmd, ok := ctr.Cmd.(*cobra.Command); ok {
			if ctrCmd.Use == cmd.Use {
				return ctr, nil
			}
		} else {
			return nil, errors.New(conf.CmdTypeAssertionFailed)
		}
	}
	return nil, errors.New(conf.CmdNotFoundInCtrSet)
}
