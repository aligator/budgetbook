package category

// Cat represents any type of financial transaction category.
type Cat struct {
	id       string
	name     string
	isInc    bool
	isCapped bool
	budget   int
}

// Helper struct for making the Cat type's fields exportable.
type exporter struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsInc    bool   `json:"is_inc"`
	IsCapped bool   `json:"capped"`
	Budget   int    `json:"budget"`
}

// Creates a new instance of Cat and returns a pointer to that instance.
// Although Tx is an exported type, it is recommended to use this factory
// since it includes some validations.
func New(name string, isInc bool, isCapped bool, budget int) *Cat {
	if isInc {
		isCapped = false
	}
	if !isCapped {
		budget = 0
	}
	id := RetrieveID(name)
	c := &Cat{
		id:       id,
		name:     name,
		isInc:    isInc,
		isCapped: isCapped,
		budget:   budget,
	}
	return c
}

// Implements Entity.MarshalJSON().
func (c *Cat) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// Implements Entity.UnmarshalJSON().
func (c *Cat) UnmarshalJSON(json []byte) error {
	return nil
}

// Implements Entity.ID().
func (c *Cat) ID() string { return c.id }

// Getter for name.
func (c *Cat) Name() string { return c.name }

// Getter for isInc.
func (c *Cat) IsInc() bool { return c.isInc }

// Getter for isCapped.
func (c *Cat) IsCapped() bool { return c.isCapped }

// Getter for budget.
func (c *Cat) Budget() int { return c.budget }

// Creates an appropriate category ID based on its name.
func RetrieveID(name string) string {
	// In this particular case the ID just corresponds to the name. However,
	// if that changes at any time, this function won't loose its validity.
	return name
}
