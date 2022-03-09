package servererrors

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type FieldError struct {
	err validator.FieldError
}

func (f FieldError) String() string {
	var sb strings.Builder

	sb.WriteString("validation failed on field '" + f.err.Field() + "'")
	sb.WriteString(", condition: " + f.err.ActualTag())

	// Print condition parameters, e.g. one_of=red blue -> { red blue }
	if f.err.Param() != "" {
		sb.WriteString(" { " + f.err.Param() + " }")
	}

	if f.err.Value() != nil && f.err.Value() != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", f.err.Value()))
	}

	return sb.String()
}

func NewFieldError(err validator.FieldError) FieldError {
	return FieldError{err: err}
}
