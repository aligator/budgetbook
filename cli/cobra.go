package cli

import (
	"budgetBookArch/intc"
	"budgetBookArch/str"
	"errors"
	"github.com/spf13/cobra"
)

type _cobra struct {
	RootCtr *container
	CtrSet  []*container
}

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
	return nil, errors.New(str.CmdTypeAssertionFailed)
}

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

func (c *_cobra) findInContainers(cmd *cobra.Command) (*container, error) {
	if rootCmd, ok := c.RootCtr.Cmd.(*cobra.Command); ok {
		if rootCmd.Use == cmd.Use {
			return c.RootCtr, nil
		}
	} else {
		return nil, errors.New(str.CmdTypeAssertionFailed)
	}
	for _, ctr := range c.CtrSet {
		if ctrCmd, ok := ctr.Cmd.(*cobra.Command); ok {
			if ctrCmd.Use == cmd.Use {
				return ctr, nil
			}
		} else {
			return nil, errors.New(str.CmdTypeAssertionFailed)
		}
	}
	return nil, errors.New(str.CmdNotFoundInCtrSet)
}
