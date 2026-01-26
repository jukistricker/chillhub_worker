
package errors

type AppError struct {
	Status  int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}
