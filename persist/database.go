package persist

import (
	"budgetBookArch/cmp"
	"budgetBookArch/conf"
	"github.com/boltdb/bolt"
)

// Database provides functions for the common CRUD operations that can be
// performed on any implementation of Entity. All further processing is up
// to the respective implementation as every database may store the entities
// in a different way.
type Database interface {
	// Opens a new database connection or file.
	Open() error
	// Returns an entity with the specified id from a given table.
	Select(id, table string) cmp.Entity
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
