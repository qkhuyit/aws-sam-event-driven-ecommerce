package errors

import "net/http"

type ErrorCode string

func (e ErrorCode) ToString() string {
	return string(e)
}

const (
	ModelInvalidErrorCode     ErrorCode = "0E00400-000"
	ResourceNotFoundErrorCode ErrorCode = "0E00400-404"
	InvalidActionErrorCode    ErrorCode = "0E00400-405"
)

var StatusCodeMap = map[ErrorCode]int{
	ModelInvalidErrorCode:     http.StatusBadRequest,
	ResourceNotFoundErrorCode: http.StatusNotFound,
	InvalidActionErrorCode:    http.StatusBadRequest,
}
