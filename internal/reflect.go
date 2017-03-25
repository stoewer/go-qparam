// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

package internal

import (
	"encoding"
	"errors"
	"reflect"
	"strconv"
)

// ParseInto parses a string into the target value.
//
// The function panics on errors.
func ParseInto(str string, target interface{}) {
	switch target := target.(type) {
	case encoding.TextUnmarshaler:
		target.UnmarshalText([]byte(str))
	default:
		value := reflect.ValueOf(target)
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}
		kind := value.Kind()

		switch kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			i, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				panic(err)
			}
			value.SetInt(i)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			u, err := strconv.ParseUint(str, 10, 64)
			if err != nil {
				panic(err)
			}
			value.SetUint(u)
		case reflect.Float32, reflect.Float64:
			f, err := strconv.ParseFloat(str, 64)
			if err != nil {
				panic(err)
			}
			value.SetFloat(f)
		case reflect.Bool:
			b, err := strconv.ParseBool(str)
			if err != nil {
				panic(err)
			}
			value.SetBool(b)
		case reflect.String:
			value.SetString(str)
		default:
			panic(errors.New("Target type must be of kind string, int, uint, float or bool"))
		}
	}
}
