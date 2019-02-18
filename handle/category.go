package handle

import (
	"budgetBookArch/cmp/category"
	"budgetBookArch/intc"
	"budgetBookArch/persist"
	"strconv"
)

type catController struct {
	db    persist.Database
	table string
}

func (c *catController) Create(cmd *intc.Command) error {
	budget, err := strconv.Atoi(cmd.P("budget"))
	if err != nil {
		budget = 0
	}
	cat := category.New(cmd.P("name"), cmd.Opt("is-inc"), cmd.Opt("is-capped"), budget)
	return c.db.Insert(cat.ID(), cat, c.table)
}

func (c *catController) DAO() persist.Database {
	return c.db
}
