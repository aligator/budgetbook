package transaction

import (
	"budgetBook/cmp"
	"budgetBook/cmp/category"
	"encoding/json"
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
	ID       string        `json:"id"`
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

// Implements Entity.MarshalJSON().
func (t *Tx) MarshalJSON() ([]byte, error) {
	return json.Marshal(&exporter{
		ID:       t.id,
		Date:     t.date,
		TxType:   t.txType,
		Category: t.category,
		Value:    t.value,
	})
}

// Implements Entity.UnmarshalJSON().
func (t *Tx) UnmarshalJSON(b []byte) error {
	exp := &exporter{}
	if err := json.Unmarshal(b, exp); err != nil {
		return err
	}
	t.id = exp.ID
	t.date = exp.Date
	t.txType = exp.TxType
	t.category = exp.Category
	t.value = exp.Value
	return nil
}

// Implements Entity.ID().
func (t *Tx) ID() string { return t.id }

// Getter for date.
func (t *Tx) Date() time.Time { return t.date }

// Getter for txType.
func (t *Tx) TxType() cmp.Type { return t.txType }

// Getter für category.
func (t *Tx) Category() *category.Cat { return t.category }

// Getter für value.
func (t *Tx) Value() int { return t.value }
