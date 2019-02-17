package persist

import "github.com/boltdb/bolt"

type _bolt struct {
	db *bolt.DB
}

func (b *_bolt) Open() error {
	return nil
}

func (b *_bolt) Select(id []byte, table string) []byte {
	return nil
}

func (b *_bolt) Insert(id, val []byte, table string) error {
	return nil
}

func (b *_bolt) Update(id, val []byte, table string) error {
	return nil
}

func (b *_bolt) Delete(id []byte, table string) error {
	return nil
}

func (b *_bolt) Close() error {
	return nil
}
