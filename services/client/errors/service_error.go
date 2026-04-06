package errors

type ServiceError struct {
	StatusCode int
	Err        error
}

func (ep ServiceError) Error() string {
	return ep.Err.Error()
}

func (ep ServiceError) Unwrap() error {
	return ep.Err
}

func NewServiceError(sc int, err error) ServiceError {
	return ServiceError{
		StatusCode: sc,
		Err:        err,
	}
}
