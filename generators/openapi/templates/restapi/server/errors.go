package <%=openApiGenPackage%>

import "fmt"

// these are errors returned by our REST API
var (
	TaskError = "task_error"
	NotRunning = "not_running"

	UnmarshalErrorCode = "unmarshall_error"
	FieldLengthErrorCode = "field_length"
	EmailErrorCode = "email"
)

type ApiError interface {
	ToApiError() Error
}

type UnmarshalError struct {
	Err error
}

type ValidationError struct {
	Code string
	Details *map[string]interface{}
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s", v.Code)
}

func (v ValidationError) ToApiError() Error {
	return Error{
		Code:    v.Code,
		Details: v.Details,
		Message: v.Error(),
	}
}

func (u UnmarshalError) Error() string {
	return u.Err.Error()
}

func (u UnmarshalError) ToApiError() Error {
	return Error{
		Code:    UnmarshalErrorCode,
		Message: u.Error(),
	}
}
