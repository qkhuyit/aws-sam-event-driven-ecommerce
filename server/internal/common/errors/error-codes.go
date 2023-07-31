package errors

import "net/http"

type ErrorCode string

const (
	ModelInvalidErrorCode     ErrorCode = "0E00400-000"
	ResourceNotFoundErrorCode ErrorCode = "0E00400-404"
)

var StatusCodeMap = map[ErrorCode]int{
	ModelInvalidErrorCode:     http.StatusBadRequest,
	ResourceNotFoundErrorCode: http.StatusNotFound,
}
