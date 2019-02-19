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
	Open() error
	Select(id, table string) cmp.Entity
	Insert(id string, e cmp.Entity, table string) error
	Update(id string, e cmp.Entity, table string) error
	Delete(id, table string) error
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
