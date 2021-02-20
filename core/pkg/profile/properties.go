package profile

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/erickmaria/kangaroo/core/pkg/handler"
	"github.com/erickmaria/kangaroo/core/pkg/utils"
)

var properties = make(map[string]interface{})

func MapSyncEnv(in interface{}, envPrefix string) error {

	var envPrefixSave string = envPrefix
	var prop, env string

	val := reflect.ValueOf(in)

	for _, key := range val.MapKeys() {

		skey := fmt.Sprint(key)
		value := val.MapIndex(key)

		envPrefix = envPrefix + skey + "_"
		env = strings.ToUpper(envPrefix[:len(envPrefix)-1])

		typ := reflect.TypeOf(value.Interface())

		if typ.Kind() == reflect.Map {

			inter := value.Interface()
			MapSyncEnv(inter, envPrefix)
			envPrefix = envPrefixSave
			continue
		}

		prop = strings.ReplaceAll(strings.ToLower(env), "_", ".")
		properties[prop] = value.Interface()

		getEnv := os.Getenv(env)

		if len(getEnv) > 0 {
			properties[prop] = reflect.ValueOf(getEnv).Interface()
		}

		envPrefix = envPrefixSave
	}

	return nil
}

func StructSyncEnv(in interface{}, envPrefix string) error {

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

				err = StructSyncEnv(valueOf.Field(i).Addr().Interface(), envPrefix)
				if handler.Validate(err) {
					return err
				}
				envPrefix = envPrefixSave
				continue
			}

			getEnv := os.Getenv(env)

			if len(getEnv) > 0 {
				err := envSync(valueOf.Field(i), getEnv)
				if handler.Validate(err) {
					return err
				}
			}

			properties[strProp] = valueOf.Field(i).Interface()
		}

		envPrefix = envPrefixSave
	}

	return nil
}

func envSync(valueOf reflect.Value, getEnv string) error {

	switch valueOf.Kind() {
	case reflect.String:
		valueOf.SetString(getEnv)
	case reflect.Int:
		toInt, err := strconv.ParseInt(getEnv, 10, 0)
		if handler.Validate(err) {
			return err
		}
		valueOf.SetInt(toInt)
	case reflect.Int8:
		toInt8, err := strconv.ParseInt(getEnv, 10, 8)
		if handler.Validate(err) {
			return err
		}
		valueOf.SetInt(toInt8)
	case reflect.Int16:
		toInt16, err := strconv.ParseInt(getEnv, 10, 16)
		if handler.Validate(err) {
			return err
		}
		valueOf.SetInt(toInt16)
	case reflect.Int32:
		toInt32, err := strconv.ParseInt(getEnv, 10, 32)
		if handler.Validate(err) {
			return err
		}
		valueOf.SetInt(toInt32)
	case reflect.Int64:
		toInt64, err := strconv.ParseInt(getEnv, 10, 64)
		if handler.Validate(err) {
			return err
		}
		valueOf.SetInt(toInt64)
	case reflect.Float32:
		toFloat32, err := strconv.ParseFloat(getEnv, 32)
		if handler.Validate(err) {
			return err
		}
		valueOf.SetFloat(toFloat32)
	case reflect.Float64:
		toFloat64, err := strconv.ParseFloat(getEnv, 64)
		if handler.Validate(err) {
			return err
		}
		valueOf.SetFloat(toFloat64)
	case reflect.Bool:
		toBool, err := strconv.ParseBool(getEnv)
		if handler.Validate(err) {
			return err
		}
		valueOf.SetBool(toBool)
	case reflect.Slice:
		strWithSpace := utils.RemoveSpace(getEnv)
		toSlice := strings.Split(strWithSpace, ",")
		sliceValue := reflect.ValueOf(toSlice)
		valueOf.Set(sliceValue)
	default:
		err := errors.New("type " + valueOf.Kind().String() + " not suported")
		return err
	}

	return nil
}

func getProfileProperties(profile Profile, profileActive string) string {
	return profile.Configs.Path + "/" + strings.Replace(profile.Configs.File, profileConfigFileKey, profileActive, 1)
}

func GetProperties(key string) interface{} {
	if len(key) == 0 {
		return properties
	}
	return properties[key]
}
