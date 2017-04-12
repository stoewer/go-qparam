// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

package internal

import (
	"errors"
	"reflect"
	"strconv"
)

// ErrUnsupportedType is the error used if the target type is not supported
var ErrUnsupportedType = errors.New("Target type must be of kind string, int, uint, float or bool")

// ParseInto parses a string into the target value.
//
// The function panics on errors.
func ParseInto(str string, target reflect.Value) {
	if unmarshal := target.MethodByName("UnmarshalText"); unmarshal.IsValid() {
		returned := unmarshal.Call([]reflect.Value{reflect.ValueOf([]byte(str))})
		if len(returned) > 0 && !returned[0].IsNil() {
			panic(errors.New("Error while calling UnmarshalText"))
		}
	} else {
		if target.Kind() == reflect.Ptr {
			target = target.Elem()
		}
		kind := target.Kind()

		switch kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			i, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				panic(err)
			}
			target.SetInt(i)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			u, err := strconv.ParseUint(str, 10, 64)
			if err != nil {
				panic(err)
			}
			target.SetUint(u)
		case reflect.Float32, reflect.Float64:
			f, err := strconv.ParseFloat(str, 64)
			if err != nil {
				panic(err)
			}
			target.SetFloat(f)
		case reflect.Bool:
			b, err := strconv.ParseBool(str)
			if err != nil {
				panic(err)
			}
			target.SetBool(b)
		case reflect.String:
			target.SetString(str)
		default:
			panic(ErrUnsupportedType)
		}
	}
}
