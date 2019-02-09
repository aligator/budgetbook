package transaction

import (
	"budgetBook/cmp"
	"budgetBook/cmp/category"
	"time"
)

// Tx represents any type of financial transaction.
type Tx struct {
	// Embedding cmp.Entity to obtain entity properties.
	cmp.Entity
	Date     time.Time     `json:"date"`
	TxType   cmp.Type      `json:"type"`
	Category *category.Cat `json:"category"`
	Value    int           `json:"value"`
}

// Creates a new instance of Tx and returns a pointer to that instance.
// Although Tx is an exported type, it is recommended to use this factory
// since it includes the creation of the entity ID.
func New(date time.Time, txType cmp.Type, category *category.Cat, value int) *Tx {
	id := RetrieveID(date)
	t := &Tx{
		Entity: cmp.Entity{
			ID: id,
		},
		Date:     date,
		TxType:   txType,
		Category: category,
		Value:    value,
	}
	return t
}

// Creates an appropriate entity ID by converting a given date into a
// corresponding string value.
func RetrieveID(date time.Time) string {
	layout := time.RFC3339
	return date.Format(layout)
}
