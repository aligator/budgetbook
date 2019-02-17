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

func (c *catController) DAO() persist.Database {
	return c.db
}

func (c *catController) Create(cmd *intc.Command) error {
	budget, err := strconv.Atoi(cmd.P("budget"))
	if err != nil {
		budget = 0
	}
	cat := category.New(cmd.P("name"), cmd.Opt("is-inc"), cmd.Opt("is-capped"), budget)
	bytes, err := cat.MarshalJSON()
	if err != nil {
		return err
	}
	id := []byte(cat.ID())
	return c.db.Insert(id, bytes, c.table)
}
