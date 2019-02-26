package handle

import (
	"budgetBook/cmp/category"
	"budgetBook/intc"
	"budgetBook/persist"
	"fmt"
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
func (c *catController) Create(cmd *intc.Command) error {
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
func (c *catController) Show(cmd *intc.Command) error {
	cats := c.db.SelectAll(c.table)
	for _, cat := range cats {
		fmt.Println(cat.ID())
	}
	return nil
}

// Implements Controller.DAO().
func (c *catController) DAO() persist.Database {
	return c.db
}
