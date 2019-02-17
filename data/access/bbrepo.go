package access

import (
	"budgetBookArch/data/persist"
	"budgetBookArch/domain/cmp"
)

type bbRepo struct {
	DB persist.Database
}

func (b *bbRepo) Find(id string) cmp.Entity {
	return nil
}

func (b *bbRepo) FindAll() []cmp.Entity {
	return nil
}

func (b *bbRepo) Insert(e cmp.Entity) error {
	return nil
}

func (b *bbRepo) Update(e cmp.Entity) error {
	return nil
}

func (b *bbRepo) Delete(e cmp.Entity) error {
	return nil
}

