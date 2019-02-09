package category

import "budgetBook/cmp"

// Cat represents any type of financial transaction category.
type Cat struct {
	// Embedding cmp.Entity to obtain entity properties.
	cmp.Entity
	Name     string `json:"name"`
	IsInc    bool   `json:"is_inc"`
	IsCapped bool   `json:"capped"`
	Budget   int    `json:"budget"`
}

// Creates a new instance of Cat and returns a pointer to that instance.
func New(name string, isInc bool, isCapped bool, budget int) *Cat {
	c := &Cat{
		Entity: cmp.Entity{
			ID: name,
		},
		Name:     name,
		IsInc:    isInc,
		IsCapped: isCapped,
		Budget:   budget,
	}
	return c
}
