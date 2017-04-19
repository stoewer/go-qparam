// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

// Package qparam provides convenient functions to read query parameters from request
// URLs. On errors the package functions create errors which provide detailed information about
// missing parameters and other failure causes.
package qparam

import (
	"errors"
	"net/url"
	"reflect"
	"strings"

	"github.com/stoewer/go-qparam/internal"
)

var (
	defaultTag    = "param"
	defaultMapper = strings.ToLower
	errNoPtr      = errors.New("Target must be a pointer")
	errNoStruct   = errors.New("Target must be a struct")
)

// ErrorMapper is an error which also contains a map of additional (named) errors
// which altogether caused the actual failure.
type ErrorMapper interface {
	error
	ErrorMap() map[string]error
}

// Option is a functional option which can be applied to a reader.
type Option func(*ParameterReader)

// Mapper is a functional option which allows to specify a custom name mapper to
// the reader.
func Mapper(mapper func(string) string) Option {
	return func(r *ParameterReader) {
		r.mapper = mapper
	}
}

// Tag is a functional option which allows to specify a custom struct tag for the
// reader.
func Tag(tag string) Option {
	return func(r *ParameterReader) {
		r.tag = tag
	}
}

// ParameterReader defines methods which can read query parameters and assign them to matching
// fields of target structs.
type ParameterReader struct {
	tag    string
	mapper func(string) string
}

// New creates a new reader which can be configured with predefined functional options.
func New(options ...Option) *ParameterReader {
	r := &ParameterReader{tag: defaultTag, mapper: defaultMapper}

	for _, opt := range options {
		opt(r)
	}

	return r
}

// ReadParams takes the provided query parameter and assigns them to the matching fields of the
// target structs.
func (r *ParameterReader) ReadParams(params url.Values, targets ...interface{}) error {
	for _, target := range targets {
		err := r.readParams(params, target)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *ParameterReader) readParams(params url.Values, target interface{}) error {
	targetVal := reflect.ValueOf(target)
	if targetVal.Kind() != reflect.Ptr {
		return errNoPtr
	}

	targetVal = targetVal.Elem()
	if targetVal.Kind() != reflect.Struct {
		return errNoStruct
	}

	targetType := reflect.TypeOf(target).Elem()
	for i := 0; i < targetType.NumField(); i++ {
		field := targetType.Field(i)

		paramName := r.paramName(field)
		paramVal, ok := params[paramName]
		if !ok || len(paramVal) == 0 {
			continue
		}

		fieldVal := targetVal.Field(i)
		r.parseSingle(paramVal[0], fieldVal)
	}

	return nil
}

func (r *ParameterReader) paramName(field reflect.StructField) string {
	name, ok := field.Tag.Lookup(r.tag)
	if !ok {
		name = r.mapper(field.Name)
	}
	return name
}

func (r *ParameterReader) parseSingle(str string, target reflect.Value) error {
	for _, parser := range internal.RegisteredCheckedParsers {
		if parser.Check(target) {
			err := parser.Parse(target, str)
			if err != nil {
				return err
			}
			return nil
		}
	}

	parser, ok := internal.RegisteredParsers[target.Kind()]
	if !ok {
		return errors.New("Not supported")
	}

	value, err := parser.Parse(str)
	if err != nil {
		return err
	}

	target.Set(value)
	return nil
}
