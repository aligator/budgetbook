package persist

type Database interface {
	Open() error
	Select(id []byte, table string) []byte
	Insert(id, val []byte, table string) error
	Update(id, val []byte, table string) error
	Delete(id []byte, table string) error
	Close() error
}

func New() Database {
	return nil
}
