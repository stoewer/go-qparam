// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

// Package qparam provides convenient functions to read query parameters from request
// URLs. On errors the package functions create errors which provide detailed information about
// missing parameters and other failure causes.
package qparam

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strings"

	"github.com/stoewer/go-qparam/internal"
)

var (
	defaultTag    = "param"
	defaultMapper = strings.ToLower

	errNoPtr        = errors.New("target must be a pointer")
	errNoStruct     = errors.New("target must be a struct")
	errNotSupported = errors.New("field has unsupported type")
)

// MultiError is an error which also contains a map of additional (named) errors
// which altogether caused the actual failure.
type MultiError map[string]error

// Error implements error for MultiError
func (e MultiError) Error() string {
	return fmt.Sprintf("%d errors occured while parsing fields", len(e))
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

// Reader defines methods which can read query parameters and assign them to matching
// fields of target structs.
type Reader struct {
	tag    string
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
func (r *Reader) Read(params url.Values, targets ...interface{}) error {
	for _, target := range targets {
		err := r.read(params, target)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Reader) read(params url.Values, target interface{}) error {
	targetVal := reflect.ValueOf(target)
	if targetVal.Kind() != reflect.Ptr {
		return errNoPtr
	}

	targetVal = targetVal.Elem()
	if targetVal.Kind() != reflect.Struct {
		return errNoStruct
	}

	it := internal.NewIterator(targetVal, r.tag, r.mapper)
	for it.HasNext() {
		name, field := it.Next()
		if values, ok := params[name]; ok && len(values) > 0 {
			checked, ok := internal.SelectCheckedParser(field)
			if ok {
				if field.Kind() == reflect.Struct {
					it.SkipStruct()
				}
				err := checked.Parse(field, values[0])
				if err != nil {
					return err
				}
				continue
			}

			parser, ok := internal.SelectParser(field)
			if !ok {
				return errNotSupported
			}

			parsed, err := parser.Parse(values[0])
			if err != nil {
				return err
			}
			field.Set(parsed)
		}
	}

	return nil
}
