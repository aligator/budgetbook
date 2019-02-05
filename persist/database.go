package persist

import (
	"github.com/boltdb/bolt"
	"os"
)

const (
	dbName   = "budgetbook"
	txTable  = "transactions"
	catTable = "categories"
)

// Database provides functions for the common CRUD operations. Since this
// interface depicts the most abstracted form of a database, it simply
// stores byte arrays.
type Database interface {
	// Initializes the DB and creates the required tables if necessary.
	Setup() error
	// Inserts an byte array value. Returns an error if insertion fails.
	Insert(id, value []byte, table string) error
	// Selects all rows from a given table.
	SelectAll(table string) [][]byte
	// Selects one specific row matching a given id.
	SelectById(table string, id []byte) []byte
}

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
		"tx":   []byte(txTable),
		"cat":  []byte(catTable),
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
	return nil
}

// Implements Database.SelectAll().
func (b *Bolt) SelectAll(table string) [][]byte {
	return nil
}

// Implements Database.SelectById().
func (b *Bolt) SelectById(table string, id []byte) []byte {
	return nil
}
