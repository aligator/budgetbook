package category

import (
	"encoding/json"
)

// cat represents any type of financial transaction category.
type cat struct {
	id       string
	name     string
	isInc    bool
	isCapped bool
	budget   int
}

// Helper struct for making the cat type's fields exportable. This is needed
// to marshal and unmarshal the entity from JSON (see below).
type exporter struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsInc    bool   `json:"is_inc"`
	IsCapped bool   `json:"capped"`
	Budget   int    `json:"budget"`
}

// Implements Entity.MarshalJSON().
// Since the json.Marshal() function can only access exported fields, the
// exporter - a identical type with exported fields - is marshaled instead
// of the actual type.
func (c *cat) MarshalJSON() ([]byte, error) {
	return json.Marshal(&exporter{
		ID:       c.id,
		Name:     c.name,
		IsInc:    c.isInc,
		IsCapped: c.isCapped,
		Budget:   c.budget,
	})
}

// Implements Entity.UnmarshalJSON().
// exporter is used for the same reason as in MarshalJSON(). The exporter
// struct will be filled and the actual type's values are read from it.
func (c *cat) UnmarshalJSON(b []byte) error {
	exp := &exporter{}
	if err := json.Unmarshal(b, exp); err != nil {
		return err
	}
	c.id = exp.ID
	c.name = exp.Name
	c.isInc = exp.IsInc
	c.isCapped = exp.IsCapped
	c.budget = exp.Budget
	return nil
}

// Implements Entity.ID().
func (c *cat) ID() string { return c.id }

// Implements Entity.ToString().
func (c *cat) ToString() string {
	s := "=== Category: " + c.name + " ===\n"
	s += "Type: "
	if c.isInc {
		s += "Incomes"
	} else {
		s += "Expenditures"
	}
	s += "\n"
	if c.budget > 0 {
		s += "Budget: " + string(c.budget) + "\n"
	}
	return s
}

// Creates a new instance of cat and returns a pointer to that instance.
// Using this factory function is required as it includes some validations.
func New(name string, isInc bool, isCapped bool, budget int) *cat {
	if isInc {
		isCapped = false
	}
	if !isCapped {
		budget = 0
	}
	id := RetrieveID(name)
	c := &cat{
		id:       id,
		name:     name,
		isInc:    isInc,
		isCapped: isCapped,
		budget:   budget,
	}
	return c
}

// Creates a new empty instance of cat and returns a pointer to that instance.
// Its main purpose is to serve as an instance to be populated with unmarshalled
// JSON data.
func Empty() *cat {
	c := &cat{}
	return c
}

// Creates an appropriate category ID based on its name.
func RetrieveID(name string) string {
	// In this particular case the ID just corresponds to the name. However,
	// if that changes at any time, this function won't loose its validity.
	return name
}
