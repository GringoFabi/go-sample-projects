package error_handler

import "fmt"

const (
	IncompleteTrainerDataErr = "provided trainer misses data"
	MissingPathParamErr = "expected path param was not provided"
	MongoFindQueryErr = "could not find document in db"
	MongoPostQueryErr = "could not save document in db"
	UnmarshalTrainerErr = "could not read provided trainer data"
)

type HandledError struct {
	Err error
	Code int
}

func (he *HandledError) Error() string {
	return fmt.Sprintf("status %d: err %v", he.Code, he.Err)
}
