package handle

import (
	"budgetBookArch/conf"
	"budgetBookArch/persist"
)

type Controller interface {
	DAO() persist.Database
}

func NewCatController() *catController {
	c := &catController{
		db: persist.New(),
		table: conf.CatTable,
	}
	return c
}
