package errors

type HTTPError struct {
	StatusCode int
	Message    string
	Cause      error
}

func (e *HTTPError) Error() string {
	if e.Cause == nil {
		return e.Message
	}

	return e.Message + " : " + e.Cause.Error()
}

func NewHttpError(err error, statusCode int, message string) error {
	return &HTTPError{
		Cause:      err,
		StatusCode: statusCode,
		Message:    message,
	}
}
