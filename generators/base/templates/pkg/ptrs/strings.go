package ptrs

import "github.com/google/uuid"

func Str(str string) *string {
	return &str
}

func UuidToStr(id uuid.UUID) *string {
	return Str(id.String())
}

func ErrToStr(err error) *string  {
	return Str(err.Error())
}
