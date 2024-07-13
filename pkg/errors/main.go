package errors

import (
	"fmt"
)

func NewNotFoundError(e string) error {
	return fmt.Errorf("%s not found", e)
}

func NewInvalidParamError(m string) error {
	return fmt.Errorf("invalid param: %s", m)
}

func NewMissingParamError(m string) error {
	return fmt.Errorf("missing param: %s", m)
}
