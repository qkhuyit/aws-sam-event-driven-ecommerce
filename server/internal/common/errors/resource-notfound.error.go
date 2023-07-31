package errors

import "fmt"

type ResourceNotfoundError struct {
	ResourceName     string
	ResourceKey      string
	ResourceKeyValue string
}

func NewResourceNotfoundError(resourceName string, resourceKey string, resourceKeyValue string) ResourceNotfoundError {
	return ResourceNotfoundError{}
}

func (m ResourceNotfoundError) Status() int {
	return StatusCodeMap[ResourceNotFoundErrorCode]
}

func (m ResourceNotfoundError) Message() string {
	return fmt.Sprintf("%s with %s = %s is not exists.", m.ResourceName, m.ResourceKey, m.ResourceKeyValue)
}

func (m ResourceNotfoundError) MessageId() string {
	return string(ResourceNotFoundErrorCode)
}

func (m ResourceNotfoundError) Error() string {
	return fmt.Sprintf("Request resource not found.")
}
