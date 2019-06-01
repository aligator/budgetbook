package app

import (
	"budgetBook/cli"
	"budgetBook/gui"
	"budgetBook/intc"
	"budgetBook/persist"
	"fmt"
	"log"
	"net/http"
)

type App interface {
	Run()
}

// a simple Writer to StdOut using fmt.Println
type StdWriter struct{}

func (StdWriter) Write(p []byte) (n int, err error) {
	return fmt.Println(p)
}

// app represents the application itself. Essentially, it consists of a root
// command next to a set of sub-commands, a CLI and a database implementation.
// While it initiates the data flow, the domain specific business logic is
// depicted in the respective command.
type app struct {
	Root     *intc.Command
	CmdSet   []*intc.Command
	Mediator cli.CLI
	DB       persist.Database
}

// Runs the application by parsing the CLI input and executing the gathered
// command which delegates the processing to the provided handler.
func (a *app) Run() {
	//defer a.DB.Close()
	exec, _ := a.Mediator.Parse()

	// Check if a handler has been assigned to the executed command.
	if exec != nil && exec.Run != nil {
		if err := exec.Run(exec, StdWriter{}); err != nil {
			fmt.Println(err)
		}
	}
}

// Creates a new instance of app and returns a pointer to that instance.
// This factory function is required since the zero value state of app is
// not usable without registering the commands in the CLI mediator.
func New() *app {
	root, cmdSet := buildCmds()
	a := &app{
		Root:     root,
		CmdSet:   cmdSet,
		Mediator: cli.New(),
		DB:       nil,
	}
	a.Mediator.Register(a.Root, a.CmdSet)
	return a
}

type webApp struct {
	Root     *intc.Command
	CmdSet   []*intc.Command
	Mediator gui.Gui
	DB       persist.Database
	Addr     string
}

func (a *webApp) Run() {
	log.Fatal(http.ListenAndServe(a.Addr, nil))
	// start server
	// show all cats
	// make actions for each command
}

func NewServer(addr string) *webApp {
	root, cmdSet := buildCmds()

	a := &webApp{
		Root:     root,
		CmdSet:   cmdSet,
		Mediator: gui.NewWeb(),
		DB:       nil,
		Addr:     addr,
	}

	a.Mediator.Register(a.Root, a.CmdSet)

	return a
}
