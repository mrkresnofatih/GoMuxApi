package utils

import (
	"errors"
)

func WrapError(message string, err error) error {
	return errors.New(message + ": " + err.Error())
}
