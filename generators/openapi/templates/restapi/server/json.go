package <%=openApiGenPackage%>

import (
	"encoding/json"
	"net/http"
)

const (
	ApplicationJson = "application/json; charset=utf-8"
	ContentType     = "Content-Type"
)

type Validatable interface {
	// Valid if json object implements this interface, we call it from UnmarshalJSON method
	Valid() ValidationError
}

// decode can be this simple to start with, but can be extended
// later to support different formats and behaviours without
// changing the interface.
func UnmarshalJSON(r *http.Request, v interface{}) ApiError {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return UnmarshalError{Err: err}
	}
	if validatable, ok := v.(Validatable); ok {
		return validatable.Valid()
	}
	return nil
}


func WriteJSONPayload(w http.ResponseWriter, payload interface{}) error {
	w.Header().Add(ContentType, ApplicationJson)
	return json.NewEncoder(w).Encode(payload)
}

func BadRequest(w http.ResponseWriter, payload interface{}) error {
	w.Header().Add(ContentType, ApplicationJson)
	w.WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(w).Encode(payload)
}

func Unauthorized(w http.ResponseWriter, payload interface{}) error {
	w.Header().Add(ContentType, ApplicationJson)
	w.WriteHeader(http.StatusUnauthorized)
	return json.NewEncoder(w).Encode(payload)
}
