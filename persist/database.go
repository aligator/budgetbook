package persist

import "budgetBookPrototype/cmp"

const (
	dbname   = "budgetbook"
	txTable  = "transactions"
	catTable = "categories"
)

// Database provides functions for the common CRUD operations that can be
// performed on the models of the cmp package.
type Database interface {
	Setup() error                            // Initializes the DB and creates the required tables if necessary.
	Insert(e cmp.Entity) error               // Inserts an entity. Returns an error if insertion fails.
	SelectAll(table string) []*cmp.Entity    // Selects all rows from a given table.
	SelectById(table, id string) *cmp.Entity // Selects one specific row matching a given id.
}

// Bolt is the most common database to use since it does not need any
// database server - it stores all entities in a file instead. Yet it is
// only one of several possible types to satisfy the Database interface.
type Bolt struct {
	rootBucket string
	txBucket   string
}

// Implements Database.Setup().
func (b *Bolt) Setup() error {
	return nil
}

// Implements Database.Insert().
func (b *Bolt) Insert(e cmp.Entity) error {
	return nil
}

// Implements Database.SelectAll().
func (b *Bolt) SelectAll(table string) []*cmp.Entity {
	return nil
}

// Implements Database.SelectById().
func (b *Bolt) SelectById(table, id string) *cmp.Entity {
	return nil
}
