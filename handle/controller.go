package handle

import (
	"budgetBookArch/persist"
)

type Controller interface {
	DAO() persist.Database
}

func NewCatController() *catController {
	c := &catController{
		db: persist.New(),
	}
	return c
}
