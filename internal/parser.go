// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

package internal

import (
	"reflect"
	"strconv"
)

// Parser is used to parse a string into its value representation. The type of the returned value
// depends on the implementation of the parser. A bool parser will turn the string "1" into 'true' but
// a int64 parser would turn the same string into 'int64(1)'.
type Parser interface {
	Parse(reflect.Value, string) error
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

// SelectParser finds a Parser that matches the kind of the provided value. If such a parser
// was found the second returned value will be true, it is false otherwise.
func SelectParser(value reflect.Value) (Parser, bool) {
	parser, ok := registeredParsers[value.Kind()]
	return parser, ok
}

type parserFunc func(reflect.Value, string) error

func (fn parserFunc) Parse(value reflect.Value, s string) error {
	return fn(value, s)
}

var intParser = parserFunc(func(value reflect.Value, s string) error {
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return err
	}
	value.Set(reflect.ValueOf(int(i)))
	return nil
})

var int8Parser = parserFunc(func(value reflect.Value, s string) error {
	i, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		return err
	}
	value.Set(reflect.ValueOf(int8(i)))
	return nil
})

var int16Parser = parserFunc(func(value reflect.Value, s string) error {
	i, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		return err
	}
	value.Set(reflect.ValueOf(int16(i)))
	return nil
})

var int32Parser = parserFunc(func(value reflect.Value, s string) error {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return err
	}
	value.Set(reflect.ValueOf(int32(i)))
	return nil
})

var int64Parser = parserFunc(func(value reflect.Value, s string) error {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	value.Set(reflect.ValueOf(i))
	return nil
})

var uintParser = parserFunc(func(value reflect.Value, s string) error {
	i, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		return err
	}
	value.Set(reflect.ValueOf(uint(i)))
	return nil
})

var uint8Parser = parserFunc(func(value reflect.Value, s string) error {
	i, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return err
	}
	value.Set(reflect.ValueOf(uint8(i)))
	return nil
})

var uint16Parser = parserFunc(func(value reflect.Value, s string) error {
	i, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return err
	}
	value.Set(reflect.ValueOf(uint16(i)))
	return nil
})

var uint32Parser = parserFunc(func(value reflect.Value, s string) error {
	i, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return err
	}
	value.Set(reflect.ValueOf(uint32(i)))
	return nil
})

var uint64Parser = parserFunc(func(value reflect.Value, s string) error {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}
	value.Set(reflect.ValueOf(i))
	return nil
})

var float32Parser = parserFunc(func(value reflect.Value, s string) error {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return err
	}
	value.Set(reflect.ValueOf(float32(f)))
	return nil
})

var float64Parser = parserFunc(func(value reflect.Value, s string) error {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	value.Set(reflect.ValueOf(f))
	return nil
})

var boolParser = parserFunc(func(value reflect.Value, s string) error {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	value.Set(reflect.ValueOf(b))
	return nil
})

var stringParser = parserFunc(func(value reflect.Value, s string) error {
	value.Set(reflect.ValueOf(s))
	return nil
})
