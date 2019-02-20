package persist

import (
	"budgetBook/cmp"
	"budgetBook/conf"
	"github.com/boltdb/bolt"
)

// Database provides functions for the common CRUD operations that can be
// performed on any implementation of Entity. As every method simply takes an
// entity next to a table name, the respective implementation doesn't know which
// concrete entity type it currently processes.
type Database interface {
	// Opens a new database connection or file. In case the database can't be
	// opened or the required tables can not be created, a corresponding error
	// will be returned.
	Open() error
	// Returns an entity with the specified id from a given table.
	Select(id, table string) cmp.Entity
	// Returns all entities stored in a given table. If there weren't found any
	// entities or the table doesn't exist, the returned slice will be empty.
	SelectAll(table string) []cmp.Entity
	// Creates a new entry in the specified table. If the entity's type does
	// not fit in the table, an error will be returned.
	Insert(id string, e cmp.Entity, table string) error
	// Updates a given entity under the same preconditions as Insert().
	Update(id string, e cmp.Entity, table string) error
	// Deletes a given entity under the same preconditions as Insert().
	Delete(id, table string) error
	// Closes the database connection or file. If the database wasn't opened,
	// an error will be returned.
	Close() error
}

// Creates a new instance of a Database implementation and returns a pointer
// to that instance. Any implementation should be an unexported type to force
// the use of this factory.
// All configuration values for the implementation are defined here.
func New() Database {
	b := &_bolt{
		db:       &bolt.DB{},
		name:     conf.DbName,
		catTable: conf.CatTable,
		txTable:  conf.TxTable,
		timeout:  conf.BoltDBTimeout,
	}
	return b
}
