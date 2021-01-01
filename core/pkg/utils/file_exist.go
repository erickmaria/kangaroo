package utils

import (
	"os"

	"github.com/erickmaria/kangaroo/core/pkg/handler"
)

func CheckFileExist(filename string) error {

	_, err := os.Stat(filename)
	if handler.Validate(err) {
		return err
	}
	if os.IsNotExist(err) {
		return err
	}

	return nil
}
