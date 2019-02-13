package category

import (
	"budgetBook/cmp"
	"budgetBook/persist"
	"github.com/pkg/errors"
)

type repository struct {
	DB persist.Database
}

func NewRepository() *repository {
	r := &repository{
		DB: &persist.Bolt{},
	}
	_ = r.DB.Setup()
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
	id := e.ID()
	if bytes, err := e.(*Cat).MarshalJSON(); err != nil {
		return r.DB.Insert([]byte(id), bytes, persist.CatTable)
	}
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

func toCat(e cmp.Entity) (*Cat, error) {
	if c, ok := e.(*Cat); ok {
		return c, nil
	}
	return nil, errors.New("given entity is not an instance of Cat")
}
