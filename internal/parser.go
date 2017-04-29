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

var registeredParsers = map[reflect.Kind]Parser{
	reflect.Int:     intParser,
	reflect.Int8:    int8Parser,
	reflect.Int16:   int16Parser,
	reflect.Int32:   int32Parser,
	reflect.Int64:   int64Parser,
	reflect.Uint:    uintParser,
	reflect.Uint8:   uint8Parser,
	reflect.Uint16:  uint16Parser,
	reflect.Uint32:  uint32Parser,
	reflect.Uint64:  uint64Parser,
	reflect.Float32: float32Parser,
	reflect.Float64: float64Parser,
	reflect.Bool:    boolParser,
	reflect.String:  stringParser,
}

func SelectParser(value reflect.Value) (Parser, bool) {
	parser, ok := registeredParsers[value.Kind()]
	return parser, ok
}

type parserFunc func(string) (reflect.Value, error)

func (fn parserFunc) Parse(s string) (reflect.Value, error) {
	return fn(s)
}

var intParser = parserFunc(func(s string) (reflect.Value, error) {
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(int(i)), nil
})

var int8Parser = parserFunc(func(s string) (reflect.Value, error) {
	i, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(int8(i)), nil
})

var int16Parser = parserFunc(func(s string) (reflect.Value, error) {
	i, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(int16(i)), nil
})

var int32Parser = parserFunc(func(s string) (reflect.Value, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(int32(i)), nil
})

var int64Parser = parserFunc(func(s string) (reflect.Value, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(i), nil
})

var uintParser = parserFunc(func(s string) (reflect.Value, error) {
	i, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(uint(i)), nil
})

var uint8Parser = parserFunc(func(s string) (reflect.Value, error) {
	i, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(uint8(i)), nil
})

var uint16Parser = parserFunc(func(s string) (reflect.Value, error) {
	i, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(uint16(i)), nil
})

var uint32Parser = parserFunc(func(s string) (reflect.Value, error) {
	i, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(uint32(i)), nil
})

var uint64Parser = parserFunc(func(s string) (reflect.Value, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(i), nil
})

var float32Parser = parserFunc(func(s string) (reflect.Value, error) {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(float32(f)), nil
})

var float64Parser = parserFunc(func(s string) (reflect.Value, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(f), nil
})

var boolParser = parserFunc(func(s string) (reflect.Value, error) {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(b), nil
})

var stringParser = parserFunc(func(s string) (reflect.Value, error) {
	return reflect.ValueOf(s), nil
})
