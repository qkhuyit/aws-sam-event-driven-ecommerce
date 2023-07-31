package errors

type ModelInvalidError struct {
	err error
}

func NewModelInvalidError(err error) ModelInvalidError {
	return ModelInvalidError{
		err: err,
	}
}

func (m ModelInvalidError) Status() int {
	return StatusCodeMap[ModelInvalidErrorCode]
}

func (m ModelInvalidError) Message() string {
	return "Request data invalid."
}

func (m ModelInvalidError) MessageId() string {
	return string(ModelInvalidErrorCode)
}

func (m ModelInvalidError) Error() string {
	return m.err.Error()
}
