package transaction

import (
	"budgetBook/cmp"
	"budgetBook/persist"
)

type repository struct {
	db persist.Database
}

func NewRepository() *repository {
	r := &repository{
		db: &persist.Bolt{},
	}
	_ = r.db.Setup()
	return r
}

// Implements cmp.repository.Find().
func (r *repository) Find(id string) cmp.Entity {
	return nil
}

// Implements cmp.repository.FindAll().
func (r *repository) FindAll() []cmp.Entity {
	return nil
}

// Implements cmp.repository.Insert().
func (r *repository) Insert(e cmp.Entity) error {
	return nil
}

// Implements cmp.repository.Update().
func (r *repository) Update(e cmp.Entity) error {
	return nil
}

// Implements cmp.repository.Delete().
func (r *repository) Delete(e cmp.Entity) error {
	return nil
}
