package gui

import (
	"budgetBook/intc"
	"net/http"
)

// a wrapper arount intc.Command which can serve the result of it via http
type webcmd struct {
	*intc.Command
	parent *intc.Command
}

func (c *webcmd) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if c.Run == nil {
		w.Write([]byte(c.Help))
	} else {
		// TODO
	}
}

func (c *webcmd) Handle() {
	http.Handle("/"+c.parent.Use+"/"+c.Use, c)
}

type web struct {
	RootCtr *container
	CtrSet  []*container
}

func (w *web) Register(root *intc.Command, cmdSet []*intc.Command) {
	for _, cmd := range cmdSet {
		webc := webcmd{cmd, root}
		webc.Handle()
	}
}
