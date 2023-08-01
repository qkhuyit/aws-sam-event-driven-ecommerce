package errors

import (
	"fmt"
	"github.com/qkhuyit/aws-sam-event-driven-ecommerce/internal/common/types"
)

func NewInvalidActionErrors(targetKey string, action string, msgFormat string) types.AppError {
	return InvalidActionError{
		TargetKey: targetKey,
		Action:    action,
	}
}

type InvalidActionError struct {
	TargetKey     string
	Action        string
	MessageFormat string
}

func (e InvalidActionError) Status() int {
	return StatusCodeMap[InvalidActionErrorCode]
}

func (e InvalidActionError) Message() string {
	return e.Error()
}

func (e InvalidActionError) MessageId() string {
	return InvalidActionErrorCode.ToString()
}

func (e InvalidActionError) Error() string {
	return fmt.Sprintf(e.MessageFormat, e.TargetKey, e.Action)
}
