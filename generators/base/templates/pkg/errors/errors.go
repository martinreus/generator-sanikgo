package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

type Unmarshal struct {
	Err error
}

type Validation struct {
	Code string
	Details *map[string]interface{}
}

type Forbidden struct {
	Code string
	Details *map[string]interface{}
}

type UniqueConstraint struct {
	Field string
	Value string
}

type Conflict struct {
	Code string
	Details *map[string]interface{}
}

func (v Validation) Error() string {
	return fmt.Sprintf("validation error: %s", v.Code)
}

func (u Unmarshal) Error() string {
	return errors.Wrap(u.Err, "unmarshal error").Error()
}

func (c UniqueConstraint) Error() string {
	return fmt.Sprintf("field '%s' with value '%s' violates unique constraint", c.Field, c.Value)
}

func (c Conflict) Error() string {
	return fmt.Sprintf("conflict")
}

func (s Forbidden) Error() string {
	return fmt.Sprintf("security exception: %s", s.Code)
}