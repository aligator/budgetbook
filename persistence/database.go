package persistence

import (
	"budgetBook/component"
)

type Database interface {
	Load() error
	Insert(entity component.Entity) error
	SelectAll(table string) []*component.Entity
	SelectById(table, id string) *component.Entity
}

type Bolt struct {
}

func (b *Bolt) Load() error {
	return nil
}

func (Bolt) Insert(entity component.Entity) error {
	return nil
}

func (Bolt) SelectAll(table string) []*component.Entity {
	return nil
}

func (Bolt) SelectById(table, id string) *component.Entity {
	return &component.Entity{}
}
