package profile

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/erickmaria/kangaroo/core/pkg/handler"
	"github.com/erickmaria/kangaroo/core/pkg/utils"
)

var properties = make(map[string]interface{})

func SyncEnv(in interface{}, envPrefix string) error {

	var env string
	var err error
	var envPrefixSave string = envPrefix

	valueOf := reflect.ValueOf(in).Elem()
	typeOf := reflect.TypeOf(in).Elem()
	for i := 0; i < valueOf.NumField(); i++ {
		tag := typeOf.Field(i).Tag.Get("yaml")

		if len(tag) > 0 {

			envPrefix = envPrefix + tag + "_"
			env = strings.ToUpper(envPrefix[:len(envPrefix)-1])
			strProp := strings.ReplaceAll(strings.ToLower(env), "_", ".")

			if valueOf.Field(i).Kind() == reflect.Struct {

				err = SyncEnv(valueOf.Field(i).Addr().Interface(), envPrefix)
				if handler.Validate(err) {
					return err
				}
				envPrefix = envPrefixSave
				continue
			}

			getEnv := os.Getenv(env)

			if len(getEnv) > 0 {

				switch valueOf.Field(i).Kind() {
				case reflect.String:
					valueOf.Field(i).SetString(getEnv)
				case reflect.Int:
					toInt, err := strconv.ParseInt(getEnv, 10, 0)
					if handler.Validate(err) {
						return err
					}
					valueOf.Field(i).SetInt(toInt)
				case reflect.Int8:
					toInt8, err := strconv.ParseInt(getEnv, 10, 8)
					if handler.Validate(err) {
						return err
					}
					valueOf.Field(i).SetInt(toInt8)
				case reflect.Int16:
					toInt16, err := strconv.ParseInt(getEnv, 10, 16)
					if handler.Validate(err) {
						return err
					}
					valueOf.Field(i).SetInt(toInt16)
				case reflect.Int32:
					toInt32, err := strconv.ParseInt(getEnv, 10, 32)
					if handler.Validate(err) {
						return err
					}
					valueOf.Field(i).SetInt(toInt32)
				case reflect.Int64:
					toInt64, err := strconv.ParseInt(getEnv, 10, 64)
					if handler.Validate(err) {
						return err
					}
					valueOf.Field(i).SetInt(toInt64)
				case reflect.Float32:
					toFloat32, err := strconv.ParseFloat(getEnv, 32)
					if handler.Validate(err) {
						return err
					}
					valueOf.Field(i).SetFloat(toFloat32)
				case reflect.Float64:
					toFloat64, err := strconv.ParseFloat(getEnv, 64)
					if handler.Validate(err) {
						return err
					}
					valueOf.Field(i).SetFloat(toFloat64)
				case reflect.Bool:
					toBool, err := strconv.ParseBool(getEnv)
					if handler.Validate(err) {
						return err
					}
					valueOf.Field(i).SetBool(toBool)
				case reflect.Slice:
					strWithSpace := utils.RemoveSpace(getEnv)
					toSlice := strings.Split(strWithSpace, ",")
					sliceValue := reflect.ValueOf(toSlice)
					valueOf.Field(i).Set(sliceValue)
				default:
					err = errors.New("type " + valueOf.Field(i).Kind().String() + " not suported")
					return err
				}

			}

			properties[strProp] = valueOf.Field(i)
		}

		envPrefix = envPrefixSave
	}

	return nil
}

func getProfileApplication(profile Profile, profileActive string) string {
	return profile.Configs.Path + "/" + strings.Replace(profile.Configs.File, profileConfigFileKey, profileActive, 1)
}

func GetProperties(key string) interface{} {
	if len(key) == 0 {
		return properties
	}
	return properties[key]
}
