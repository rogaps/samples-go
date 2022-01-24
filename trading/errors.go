package trading

type NonRetryableError struct {
	message string
}

func NewNonRetryableError(msg string) *NonRetryableError {
	return &NonRetryableError{
		message: msg,
	}
}

func (e NonRetryableError) Error() string {
	return e.message
}
