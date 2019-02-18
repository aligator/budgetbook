package persist

import (
	"budgetBookArch/cmp"
	"budgetBookArch/conf"
	"github.com/boltdb/bolt"
)

type Database interface {
	Open() error
	Select(id, table string) cmp.Entity
	Insert(id string, e cmp.Entity, table string) error
	Update(id string, e cmp.Entity, table string) error
	Delete(id, table string) error
	Close() error
}

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
