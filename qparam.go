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
	errNoPtr      = errors.New("Target must be a pointer")
	errNoStruct   = errors.New("Target must be a struct")
	errMultiple   = errors.New("Parameter has multiple values but target is not a slice")
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
		targetType := reflect.TypeOf(target)
		if targetType.Kind() != reflect.Ptr {
			return errNoPtr
		}

		targetType = targetType.Elem()
		if targetType.Kind() != reflect.Struct {
			return errNoStruct
		}

		targetVal := reflect.ValueOf(target).Elem()
		for i := 0; i < targetType.NumField(); i++ {
			field := targetType.Field(i)

			paramName, ok := field.Tag.Lookup(r.tag)
			if !ok {
				paramName = r.mapper(field.Name)
			}

			paramVal, ok := params[paramName]
			if !ok {
				continue
			}
			if len(paramVal) == 0 {
				continue
			}

			fieldVal := targetVal.Field(i)
			switch fieldVal.Kind() {
			case reflect.Slice:
				return fmt.Errorf("Field type Slice is not supported yet")
			default:
				if len(paramVal) > 1 {
					return errMultiple
				}
				internal.ParseInto(paramVal[0], fieldVal)
			}
		}
	}
	return nil
}
