package constants

import "fmt"


var (
	ErrIncorrectInputData = fmt.Errorf("unable to parse input data")
	ErrEventNotFound = fmt.Errorf("event was not found")
	ErrNoRow = fmt.Errorf("sql: no rows in result set")
	ErrUserIDNotFound = fmt.Errorf("user with this ID was not found in the table")
)