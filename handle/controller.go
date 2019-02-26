package handle

import (
	"budgetBook/conf"
	"budgetBook/persist"
)

// Controller is responsible for delegating the command the user has entered
// and therefore handles commands. This means that the Controller's methods take
// interchangeable commands as arguments and processes them. Most Controller
// implementations hold an Data Access Object to delegate database operations.
type Controller interface {
	// Returns a Controller's Data Access Object which can be any implementation
	// of Database. In case the Controller doesn't hold any DAO, it returns nil.
	DAO() persist.Database
}

// Creates a new instance of catController and returns a pointer to that instance.
// This factory chooses the Database implementation being used as DAO.
func NewCatController() *catController {
	db, err := persist.New()
	if err != nil {
		panic(err)
	}
	c := &catController{
		db:    db,
		table: conf.CatTable,
	}
	return c
}
