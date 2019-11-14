// Copyright (c) 2017, A. Stoewer <adrian@stoewer.me>
// All rights reserved.

package internal

import (
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

// Parser is used to parse a string and assign it to the provided value.
type Parser interface {
	Parse(reflect.Value, string) error
}

// CheckedParser is a parser that has a method (Check) which can be used to determine whether the
// parser can be applied to a certain value. It is recommended to call Check prior to calling Parse.
type CheckedParser interface {
	Parser
	Check(reflect.Value) bool
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

var registeredCheckedParsers = []CheckedParser{
	textParser{},
}

// FindParser finds a Parser that matches the provided value. If such a parser
// was found the second returned value will be true, it is false otherwise.
func FindParser(value reflect.Value) (Parser, bool) {
	for _, parser := range registeredCheckedParsers {
		if parser.Check(value) {
			return parser, true
		}
	}

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

type textParser struct{}

func (p textParser) Check(value reflect.Value) bool {
	return p.method(value).IsValid()
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
