package transaction

import "budgetBook/cmp"

type Repository struct {}

func (r *Repository) Find(id string) *cmp.Entity {
	return nil
}

func (r *Repository) FindAll() []*cmp.Entity {
	return nil
}

func (r *Repository) Insert(e *cmp.Entity) error {
	return nil
}

func (r *Repository) Update(e *cmp.Entity) error {
	return nil
}

func (r *Repository) Delete(e *cmp.Entity) error {
	return nil
}
