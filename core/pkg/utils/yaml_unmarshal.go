package utils

import (
	"errors"
	"io/ioutil"
	"reflect"

	"github.com/erickmaria/kangaroo/core/pkg/handler"
	"gopkg.in/yaml.v2"
)

func YamlUnmarshal(in interface{}, filename string) error {

	if reflect.ValueOf(in).Kind() != reflect.Ptr {
		return errors.New("value passed not is a pointer")
	}

	profileByte, ReaFileError := ioutil.ReadFile(filename)

	if handler.Validate(ReaFileError) {
		return ReaFileError
	}
	YamlUnmarshalError := yaml.Unmarshal(profileByte, in)
	if handler.Validate(YamlUnmarshalError) {
		return YamlUnmarshalError
	}

	return nil

}
