package handle

import (
	"budgetBookArch/cmp/category"
	"budgetBookArch/intc"
	"budgetBookArch/persist"
	"strconv"
)

type catController struct {
	db persist.Database
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
	_, err = cat.MarshalJSON()
	if err != nil {
		return err
	}
	return nil
}
