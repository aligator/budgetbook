package transaction

import "budgetBook/cmp"

type Repository struct {}

// Implements cmp.Repository.Find().
func (r *Repository) Find(id string) *cmp.Entity {
	return nil
}

// Implements cmp.Repository.FindAll().
func (r *Repository) FindAll() []*cmp.Entity {
	return nil
}

// Implements cmp.Repository.Insert().
func (r *Repository) Insert(e *cmp.Entity) error {
	return nil
}

// Implements cmp.Repository.Update().
func (r *Repository) Update(e *cmp.Entity) error {
	return nil
}

// Implements cmp.Repository.Delete().
func (r *Repository) Delete(e *cmp.Entity) error {
	return nil
}
