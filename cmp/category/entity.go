package category

import "encoding/json"

type cat struct {
	id       string
	name     string
	isInc    bool
	isCapped bool
	budget   int
}

type exporter struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsInc    bool   `json:"is_inc"`
	IsCapped bool   `json:"capped"`
	Budget   int    `json:"budget"`
}

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
func RetrieveID(name string) string {
	return name
}

func (c *cat) MarshalJSON() ([]byte, error) {
	return json.Marshal(&exporter{
		ID:       c.id,
		Name:     c.name,
		IsInc:    c.isInc,
		IsCapped: c.isCapped,
		Budget:   c.budget,
	})
}

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

func (c *cat) ID() string { return c.id }

func (c *cat) Name() string { return c.name }

func (c *cat) IsInc() bool { return c.isInc }

func (c *cat) IsCapped() bool { return c.isCapped }

func (c *cat) Budget() int { return c.budget }
