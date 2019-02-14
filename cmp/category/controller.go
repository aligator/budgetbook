package category

import (
	"budgetBook/intc"
	"fmt"
	"strconv"
)

func Create(cmd *intc.Command) error {
	fmt.Println("Creating new Category")
	budget, err := strconv.Atoi(cmd.Params["budget"].S)
	if err != nil {
		return err
	}
	cat := New(cmd.Params["name"].S, cmd.Options["is-inc"].S, cmd.Options["is-capped"].S, budget)
	fmt.Println("created '", cmd.Params["name"].S, "'!")
	return NewRepository().Insert(cat)
}
