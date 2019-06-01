package handle

import (
	"budgetBook/cmp/category"
	"budgetBook/intc"
	"budgetBook/persist"
	"io"
	"strconv"
)

// catController provides any methods (respectively command handlers) that
// are associated with category.cat and its related types.
type catController struct {
	// db represents the DAO that is prescribed by the Controller interface.
	db persist.Database
	// The table that is mainly used by the controller.
	table string
}

// Creates a new instance of category.cat out of a interchangeable command
// and stores that instance in the corresponding database that is lodged
// in the DAO.
func (c *catController) Create(cmd *intc.Command, w io.Writer) error {
	budget, err := strconv.Atoi(cmd.P("budget"))
	// If the budget parameter is not a valid integer value, it will be set
	// to its zero value. This implies a budget-less, uncapped category.
	if err != nil {
		budget = 0
	}
	cat := category.New(cmd.P("name"), cmd.Opt("is-inc"), cmd.Opt("is-capped"), budget)
	return c.db.Insert(cat.ID(), cat, c.table)
}

// Retrieves all categories from the database and prints them.
func (c *catController) Show(cmd *intc.Command, w io.Writer) error {
	catBytes := c.db.SelectAll(c.table)
	for _, bytes := range catBytes {
		cat := category.Empty()
		if err := cat.UnmarshalJSON(bytes); err != nil {
			return err
		}
		w.Write([]byte(cat.ToString()))
	}
	return nil
}

// Implements Controller.DAO().
func (c *catController) DAO() persist.Database {
	return c.db
}
