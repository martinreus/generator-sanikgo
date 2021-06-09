package <%=openApiGenPackage%>

import (
	"encoding/json"
	"net/http"
	"<%=moduleName%>/pkg/errors"
)

const (
	ApplicationJson = "application/json; charset=utf-8"
	ContentType     = "Content-Type"
)

type Validatable interface {
	// Valid if json object implements this interface, we call it from unmarshalJSON method
	Valid() errors.Validation
}

// decode can be this simple to start with, but can be extended
// later to support different formats and behaviours without
// changing the interface.
func unmarshalJSON(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return errors.Unmarshal{Err: err}
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

func WriteError(w http.ResponseWriter, err error) error {
	w.Header().Add(ContentType, ApplicationJson)
	apiErr, httpStatusCode := convertErrors(err)
	w.WriteHeader(httpStatusCode)
	return json.NewEncoder(w).Encode(apiErr)
}
