// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

package internal

import (
	"reflect"

	"github.com/pkg/errors"
)

// CheckedParser is a parser that has a method (Check) which can be used to determine whether the
// parser can be applied to a certain value. It is recommended to call Check prior to calling Parse.
type CheckedParser interface {
	Check(reflect.Value) bool
	Parse(reflect.Value, string) error
}

var registeredCheckedParsers = []CheckedParser{
	textParser{},
}

// SelectCheckedParser finds a registered CheckedParser which passes its check for the provided value.
// If such a parser was found the second return value will be true, it is false otherwise.
func SelectCheckedParser(value reflect.Value) (CheckedParser, bool) {
	for _, parser := range registeredCheckedParsers {
		if parser.Check(value) {
			return parser, true
		}
	}
	return nil, false
}

type textParser struct{}

func (p textParser) Check(value reflect.Value) bool {
	unmarshal := p.method(value)
	if !unmarshal.IsValid() {
		return false
	}
	return true
}

func (p textParser) Parse(value reflect.Value, s string) error {
	unmarshal := p.method(value)
	if !unmarshal.IsValid() {
		return errors.New("method UnmarshalText not available")
	}

	returned := unmarshal.Call([]reflect.Value{reflect.ValueOf([]byte(s))})
	if len(returned) > 0 && !returned[0].IsNil() {
		return errors.Errorf("%s", returned[0])
	}

	return nil
}

func (p textParser) method(value reflect.Value) reflect.Value {
	m := value.MethodByName("UnmarshalText")
	if !m.IsValid() {
		if value.Kind() == reflect.Ptr {
			m = value.Elem().MethodByName("UnmarshalText")
		} else {
			if value.CanAddr() {
				m = value.Addr().MethodByName("UnmarshalText")
			}
		}
	}
	return m
}
