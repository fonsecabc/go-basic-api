package validations

import (
	"github.com/fonsecabc/go-basic-api/pkg/errors"
	"github.com/fonsecabc/go-basic-api/pkg/value_objects"
)

func ValidateID(id value_objects.ID) error {
	if id.String() == "" {
		return errors.NewMissingParamError("id")
	}

	if _, err := value_objects.ParseID(id.String()); err != nil {
		return errors.NewInvalidParamError("id")
	}

	return nil
}
