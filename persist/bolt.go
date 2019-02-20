package persist

import (
	"budgetBook/cmp"
	"github.com/boltdb/bolt"
	"os"
	"time"
)

// _bolt is a wrapper for BoltDB, the most common database to use as it does not
// need any database server - it stores all entities in a key-value-file instead.
// Yet it is only one of several possible types to satisfy the Database interface.
type _bolt struct {
	db       *bolt.DB
	// The database name will also be used as the identifier of the root bucket.
	name     string
	catTable string
	txTable  string
	timeout  time.Duration
}

// Implements Database.Open().
// Tries to open the bolt file and creates the required buckets. If the bolt file
// doesn't exist, it will be created with the corresponding file mode (see below).
// In case one or several buckets can't be created, an error will be returned.
func (b *_bolt) Open() error {
	var err error
	var mode os.FileMode = 0600

	b.db, err = bolt.Open(b.name, mode, &bolt.Options{Timeout: b.timeout})
	if err != nil {
		return err
	}
	create := func(btx *bolt.Tx) error {
		// The root bucket rather represents the database itself than a specific
		// table. Any tables are depicted as child buckets of root.
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

// Implements Database.Select().
func (b *_bolt) Select(id, table string) cmp.Entity {
	return nil
}

// Implements Database.Insert().
func (b *_bolt) Insert(id string, e cmp.Entity, table string) error {
	return nil
}

// Implements Database.Update().
func (b *_bolt) Update(id string, e cmp.Entity, table string) error {
	return nil
}

// Implements Database.Delete().
func (b *_bolt) Delete(id, table string) error {
	return nil
}

// Implements Database.Close().
func (b *_bolt) Close() error {
	return b.db.Close()
}
