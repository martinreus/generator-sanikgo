package <%=openApiGenPackage%>

import (
	"net/http"
	"<%=moduleName%>/pkg/errors"
	"<%=moduleName%>/pkg/ptrs"
)

// these are errors returned by our REST API
const (
	TaskError  = "task_error"
	NotRunning = "not_running"

	UniqueConstraint = "unique_constraint"
	UnmarshalErrorCode   = "unmarshall_error"
	Conflict             = "conflict"
	Unexpected           = "unexpected"
	Forbidden            = "forbidden"
)

// convertErrors converts a service error into an Api Error and an http status code
func convertErrors(err error) (Error, int) {
	switch err.(type) {
	case errors.Validation:
		valError := err.(errors.Validation)
		return Error{
			Code:    valError.Code,
			Details: valError.Details,
			Message: ptrs.ErrToStr(valError),
		}, http.StatusBadRequest
	case errors.Conflict:
		cErr := err.(errors.Conflict)
		return Error{
			Code:    Conflict,
			Details: cErr.Details,
			Message: ptrs.ErrToStr(cErr),
		}, http.StatusConflict
	case errors.Forbidden:
		fErr := err.(errors.Forbidden)
		return Error{
			Code:    Forbidden,
			Details: fErr.Details,
			Message: ptrs.ErrToStr(fErr),
		}, http.StatusForbidden
	case errors.Unmarshal:
		uErr := err.(errors.Unmarshal)
		return Error{
			Code:    UnmarshalErrorCode,
			Message: ptrs.ErrToStr(uErr),
		}, http.StatusBadRequest
	case errors.UniqueConstraint:
		uCErr := err.(errors.UniqueConstraint)
		return Error{
			Code:    UniqueConstraint,
			Message: ptrs.ErrToStr(uCErr),
			Details: &map[string]interface{}{
				"field": uCErr.Field,
				"value": uCErr.Value,
			},
		}, http.StatusConflict
	default:
		return Error{
			Code:    Unexpected,
			Message: ptrs.Str(err.Error()),
		}, http.StatusInternalServerError
	}
}
