package persist

import (
	"errors"
	"github.com/boltdb/bolt"
	"os"
)

// Bolt is the most common database to use since it does not need any
// database server - it stores all entities in a file instead. Yet it is
// only one of several possible types to satisfy the Database interface.
type Bolt struct {
	db      *bolt.DB
	buckets map[string][]byte
}

// Implements Database.Setup().
func (b *Bolt) Setup() error {
	var err error
	var mode os.FileMode = 0600

	if b.db, err = bolt.Open(dbName, mode, &bolt.Options{Timeout: 400}); err != nil {
		return err
	}
	b.buckets = map[string][]byte{
		"root": []byte("db"),
		"tx":   []byte(TxTable),
		"cat":  []byte(CatTable),
	}
	create := func(btx *bolt.Tx) error {
		root, err := btx.CreateBucketIfNotExists(b.buckets["root"])
		if err != nil {
			return err
		}
		_, err = root.CreateBucketIfNotExists(b.buckets["tx"])
		_, err = root.CreateBucketIfNotExists(b.buckets["cat"])
		return err
	}
	return b.db.Update(create)
}

// Implements Database.Insert().
func (b *Bolt) Insert(id, value []byte, table string) error {
	return b.db.Update(func(btx *bolt.Tx) error {
		bucket := btx.Bucket(b.buckets[table])
		if bucket != nil {
			return bucket.Put(id, value)
		}
		return errors.New("specified table not found")
	})
}

// Implements Database.SelectAll().
func (b *Bolt) SelectAll(table string) [][]byte {
	var r [][]byte
	selectAll := func(btx *bolt.Tx) error {
		bucket := btx.Bucket(b.buckets[table])
		if bucket != nil {
			bucket.ForEach(func(id, value []byte) error {
				r = append(r, value)
				return nil
			})
		}
		return nil
	}
	_ = b.db.View(selectAll)
	return r
}

// Implements Database.SelectById().
func (b *Bolt) SelectById(table string, id []byte) []byte {
	var r []byte
	selectById := func(btx *bolt.Tx) error {
		bucket := btx.Bucket(b.buckets[table])
		if bucket != nil {
			r = bucket.Get(id)
		}
		return nil
	}
	_ = b.db.View(selectById)
	return r
}
