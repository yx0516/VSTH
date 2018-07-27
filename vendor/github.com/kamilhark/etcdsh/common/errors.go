package common

func NewStringError(message string) error {
	return &stringError{message}
}

type stringError struct {
	s string
}

func (e *stringError) Error() string {
	return e.s
}
