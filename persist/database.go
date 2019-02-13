package persist

const (
	dbName   = "budgetbook"
	TxTable  = "transactions"
	CatTable = "categories"
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
