package handle

import (
	"budgetBookArch/intc"
	"budgetBookArch/access"
	"budgetBookArch/cmp/category"
	"strconv"
)

type catController struct {
	Repo access.Repository
}

func (c *catController) DAO() access.Repository {
	return c.Repo
}

func (c *catController) Create(cmd *intc.Command) error {
	budget, err := strconv.Atoi(cmd.P("budget"))
	if err != nil {
		budget = 0
	}
	cat := category.New(cmd.P("name"), cmd.Opt("is-inc"), cmd.Opt("is-capped"), budget)
	return c.Repo.Insert(cat)
}
