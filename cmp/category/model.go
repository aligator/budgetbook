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
