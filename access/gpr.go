package access

import (
	"budgetBookArch/cmp"
	"budgetBookArch/persist"
)

type gpr struct {
	DB persist.Database
}

func (b *gpr) Find(id string) cmp.Entity {
	return nil
}

func (b *gpr) FindAll() []cmp.Entity {
	return nil
}

func (b *gpr) Insert(e cmp.Entity) error {
	return nil
}

func (b *gpr) Update(e cmp.Entity) error {
	return nil
}

func (b *gpr) Delete(e cmp.Entity) error {
	return nil
}

