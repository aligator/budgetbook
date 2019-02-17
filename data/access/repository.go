package access

import (
	"budgetBookArch/data/persist"
	"budgetBookArch/domain/cmp"
)

type Repository interface {
	Find(id string) cmp.Entity
	FindAll() []cmp.Entity
	Insert(e cmp.Entity) error
	Update(e cmp.Entity) error
	Delete(e cmp.Entity) error
}

func New() Repository {
	b := &bbRepo{
		DB: persist.New(),
	}
	return b
}
