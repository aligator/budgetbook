package persist

import "github.com/boltdb/bolt"

type Database interface {
	Open() error
	Select(id []byte, table string) []byte
	Insert(id, val []byte, table string) error
	Update(id, val []byte, table string) error
	Delete(id []byte, table string) error
	Close() error
}

func New() Database {
	b := &_bolt{
		db: &bolt.DB{},
	}
	return b
}
