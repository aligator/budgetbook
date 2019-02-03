package cli

import (
	"budgetBookProto/intc"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"reflect"
)

type Parser struct {
	RootCmd        *cmdWrap
	Cmds           []*cmdWrap
	TfHandler      func(cmd *intc.Command) *cmdWrap
	ExecuteHandler func(rootCmd *cmdWrap) *intc.Command
}

type cmdWrap struct {
	Cmd interface{}
}

func (p *Parser) Transform(rootCmd *intc.Command, cmds []*intc.Command) {
	p.RootCmd = &cmdWrap{
		Cmd: p.TfHandler(rootCmd),
	}
	for _, intcCmd := range cmds {
		wrapper := &cmdWrap{
			Cmd: p.TfHandler(intcCmd),
		}
		p.Cmds = append(p.Cmds, wrapper)
	}
}

func (p *Parser) Execute() *intc.Command {
	return p.ExecuteHandler(p.RootCmd)
}

var Cobra = &Parser{
	RootCmd: &cmdWrap{
		&cobra.Command{},
	},
	TfHandler: func(cmd *intc.Command) *cmdWrap {
		cobraCmd := &cobra.Command{
			Use: cmd.Use,
		}
		for _, f := range cmd.Flags {
			cobraCmd.Flags().AddFlag(&pflag.Flag{
				Name:      f.Name,
				Shorthand: f.Shorthand,
				DefValue:  f.DefValue,
			})
		}
		return &cmdWrap{
			Cmd: cobraCmd,
		}
	},
	ExecuteHandler: func(rootCmd *cmdWrap) *intc.Command {
		root := rootCmd.Cmd
		cmd, _ := root.(*cobra.Command)
		fmt.Println(reflect.TypeOf(root), reflect.TypeOf(cmd))
		if true {
			execCmd, _ := cmd.ExecuteC()
			fmt.Println(execCmd.Use)
		}
		fmt.Println("ExecuteSchiefgelaufen")
		return nil
	},
}
