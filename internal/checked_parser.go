// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

package internal

import (
	"errors"
	"reflect"
)

type CheckedParser interface {
	Check(reflect.Value) bool
	Parse(reflect.Value, string) error
}

var RegisteredCheckedParsers = []CheckedParser{
	&TextParser{},
}

type TextParser struct{}

func (p *TextParser) Check(val reflect.Value) bool {
	unmarshal := val.MethodByName("UnmarshalText")
	if !unmarshal.IsValid() {
		return false
	}
	return true
}

func (p *TextParser) Parse(val reflect.Value, s string) error {
	unmarshal := val.MethodByName("UnmarshalText")
	if !unmarshal.IsValid() {
		return errors.New("Method UnmarshalText not available")
	}

	returned := unmarshal.Call([]reflect.Value{reflect.ValueOf([]byte(s))})
	if len(returned) > 0 && !returned[0].IsNil() {
		return errors.New("Error while calling UnmarshalText")
	}

	return nil
}
