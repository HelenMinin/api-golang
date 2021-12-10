package pkg

type ResponseError struct {
	Cause   string
	Message string
}

func NewResponseError(cause string, err error) *ResponseError {
	return &ResponseError{
		Cause:   cause,
		Message: err.Error(),
	}
}
