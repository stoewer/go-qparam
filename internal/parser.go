// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

package internal

import (
	"reflect"
	"strconv"
)

type Parser interface {
	Parse(string) (reflect.Value, error)
}

type ParserFunc func(string) (reflect.Value, error)

func (fn ParserFunc) Parse(s string) (reflect.Value, error) {
	return fn(s)
}

var RegisteredParsers = map[reflect.Kind]Parser{
	reflect.Int:     IntParser,
	reflect.Int8:    IntParser,
	reflect.Int16:   IntParser,
	reflect.Int32:   IntParser,
	reflect.Int64:   IntParser,
	reflect.Uint:    UintParser,
	reflect.Uint8:   UintParser,
	reflect.Uint16:  UintParser,
	reflect.Uint32:  UintParser,
	reflect.Uint64:  UintParser,
	reflect.Float32: FloatParser,
	reflect.Float64: FloatParser,
	reflect.Bool:    BoolParser,
	reflect.String:  StringParser,
}

var IntParser = ParserFunc(func(s string) (reflect.Value, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(i), nil
})

var UintParser = ParserFunc(func(s string) (reflect.Value, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(i), nil
})

var FloatParser = ParserFunc(func(s string) (reflect.Value, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(f), nil
})

var BoolParser = ParserFunc(func(s string) (reflect.Value, error) {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(b), nil
})

var StringParser = ParserFunc(func(s string) (reflect.Value, error) {
	return reflect.ValueOf(s), nil
})
