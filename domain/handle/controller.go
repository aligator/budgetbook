package handle

import "budgetBookArch/data/access"

type Controller interface {
	DAO() access.Repository
}

func NewCatController() *catController {
	c := &catController{
		Repo: access.New(),
	}
	return c
}
