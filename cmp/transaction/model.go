package transaction

import (
	"budgetBook/cmp"
	"budgetBook/cmp/category"
	"time"
)

// Tx represents any type of financial transaction.
type Tx struct {
	id       string
	date     time.Time
	txType   cmp.Type
	category *category.Cat
	value    int
}

// Helper struct for making the Cat type's fields exportable.
type exporter struct {
	ID       string       `json:"id"`
	Date     time.Time    `json:"date"`
	TxType   cmp.Type     `json:"type"`
	Category category.Cat `json:"category"`
	Value    int          `json:"value"`
}

// Creates a new instance of Tx and returns a pointer to that instance.
// Although Tx is an exported type, it is recommended to use this factory
// since it includes the creation of the entity ID.
func New(date time.Time, txType cmp.Type, category *category.Cat, value int) *Tx {
	id := RetrieveID(date)
	t := &Tx{
		id:       id,
		date:     date,
		txType:   txType,
		category: category,
		value:    value,
	}
	return t
}

// Creates an appropriate transaction ID by converting a given date into
// a corresponding string value.
func RetrieveID(date time.Time) string {
	layout := time.RFC3339
	return date.Format(layout)
}
