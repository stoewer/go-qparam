// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

// Package qparam provides convenient functions to read query parameters from request
// URLs. On errors the package functions create errors which provide detailed information about
// missing parameters and other failure causes.
package qparam

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/stoewer/go-qparam/internal"
)

var (
	defaultTag    = "param"
	defaultMapper = strings.ToLower
)

// MultiError is an error which also contains a map of additional (named) errors
// which altogether caused the actual failure.
type MultiError interface {
	error
	ErrorMap() map[string]error
}

// implementation of MultiError
type multiError map[string]error

// Error returns a string summarizing all errors
func (err multiError) Error() string {
	return fmt.Sprintf("%d errors occured while reading parameters", len(err))
}

// ErrorMap returns all field names with their respective errors
func (err multiError) ErrorMap() map[string]error {
	return err
}

// Option is a functional option which can be applied to a reader.
type Option func(*Reader)

// Mapper is a functional option which allows to specify a custom name mapper to
// the reader.
func Mapper(mapper func(string) string) Option {
	return func(r *Reader) {
		r.mapper = mapper
	}
}

// Tag is a functional option which allows to specify a custom struct tag for the
// reader.
func Tag(tag string) Option {
	return func(r *Reader) {
		r.tag = tag
	}
}

// Strict is a functional option used to define whether the reader runs in struct
// mode or not. In strict mode all parsed values must have an equivalent target field.
// If the strict rule is violated the Reader returns an error.
func Strict(strict bool) Option {
	return func(r *Reader) {
		r.strict = strict
	}
}

// Reader defines methods which can read query parameters and assign them to matching
// fields of target structs.
type Reader struct {
	tag    string
	strict bool
	mapper func(string) string
}

// New creates a new reader which can be configured with predefined functional options.
func New(options ...Option) *Reader {
	r := &Reader{tag: defaultTag, mapper: defaultMapper}

	for _, opt := range options {
		opt(r)
	}

	return r
}

// Read takes the provided query parameter and assigns them to the matching fields of the
// target structs.
//
// If an error occurs while parsing values for struct fields, the returned error probably
// implements the interface MultiError. In that case specific errors for each failed field
// can be obtained from the error.
func (r *Reader) Read(params url.Values, targets ...interface{}) error {
	var processed map[string]struct{}
	if r.strict {
		processed = make(map[string]struct{})
	}

	fieldErrors := multiError{}
	for _, target := range targets {
		targetVal := reflect.ValueOf(target)
		if targetVal.Kind() != reflect.Ptr {
			return errors.New("target must be a pointer")
		}

		targetVal = targetVal.Elem()
		if targetVal.Kind() != reflect.Struct {
			return errors.New("target must be a struct")
		}

		it := internal.NewIterator(targetVal, r.tag, r.mapper)
		for it.HasNext() {
			name, field := it.Next()
			if values, ok := params[name]; ok && len(values) > 0 {
				var err error

				if field.Kind() == reflect.Slice {
					err = r.readSlice(values, field)
				} else {
					err = r.readSingle(values, field, it)
				}

				if err != nil {
					fieldErrors[name] = err
				}

				if r.strict {
					processed[name] = struct{}{}
				}
			}
		}
	}

	if r.strict {
		for name := range params {
			if _, ok := processed[name]; !ok {
				fieldErrors[name] = errors.New("unknown parameter name")
			}
		}
	}

	if len(fieldErrors) > 0 {
		return fieldErrors
	}

	return nil
}

func (r *Reader) readSingle(values []string, field reflect.Value, it *internal.Iterator) error {
	if len(values) > 1 {
		return errors.New("multiple values for single value parameter")
	}

	checked, ok := internal.SelectCheckedParser(field)
	if ok {
		if field.Kind() == reflect.Struct {
			it.SkipStruct()
		}
		return checked.Parse(field, values[0])
	}

	parser, ok := internal.SelectParser(field)
	if !ok {
		return errors.New("target field type is not supported")
	}

	parsed, err := parser.Parse(values[0])
	if err != nil {
		return err
	}

	field.Set(parsed)
	return nil
}

func (r *Reader) readSlice(values []string, slice reflect.Value) error {
	slice.Set(reflect.MakeSlice(slice.Type(), len(values), len(values)))

	first := slice.Index(0)
	isPtr := first.Kind() == reflect.Ptr
	if isPtr {
		for i := 0; i < slice.Len(); i++ {
			slice.Index(i).Set(reflect.New(slice.Index(i).Type().Elem()))
		}
		first = first.Elem()
	}

	checked, ok := internal.SelectCheckedParser(first)
	if ok {
		for i, value := range values {
			err := checked.Parse(slice.Index(i), value)
			if err != nil {
				return err
			}
		}
		return nil
	}

	parser, ok := internal.SelectParser(first)
	if !ok {
		return errors.New("target field type is not supported")
	}

	for i, value := range values {
		parsed, err := parser.Parse(value)
		if err != nil {
			return err
		}
		if isPtr {
			slice.Index(i).Elem().Set(parsed)
		} else {
			slice.Index(i).Set(parsed)
		}
	}
	return nil
}
