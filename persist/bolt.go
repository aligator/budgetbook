package persist

import (
	"budgetBookArch/cmp"
	"github.com/boltdb/bolt"
	"os"
	"time"
)

type _bolt struct {
	db       *bolt.DB
	name     string
	catTable string
	txTable  string
	timeout  time.Duration
}

func (b *_bolt) Open() error {
	var err error
	var mode os.FileMode = 0600

	b.db, err = bolt.Open(b.name, mode, &bolt.Options{Timeout: b.timeout})
	if err != nil {
		return err
	}
	create := func(btx *bolt.Tx) error {
		root, err := btx.CreateBucketIfNotExists([]byte(b.name))
		if err != nil {
			return err
		}
		_, err = root.CreateBucketIfNotExists([]byte(b.catTable))
		_, err = root.CreateBucketIfNotExists([]byte(b.txTable))
		return err
	}
	return b.db.Update(create)
}

func (b *_bolt) Select(id, table string) cmp.Entity {
	return nil
}

func (b *_bolt) Insert(id string, e cmp.Entity, table string) error {
	return nil
}

func (b *_bolt) Update(id string, e cmp.Entity, table string) error {
	return nil
}

func (b *_bolt) Delete(id, table string) error {
	return nil
}

func (b *_bolt) Close() error {
	return b.db.Close()
}
