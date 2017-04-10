// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

// Package qparam provides convenient functions to read query parameters from request
// URLs. On errors the package functions create errors which provide detailed information about
// missing parameters and other failure causes.
package qparam

import (
	"errors"
	"fmt"
	"github.com/stoewer/go-qparam/internal"
	"net/url"
	"reflect"
	"strings"
)

var (
	defaultTag    = "param"
	defaultMapper = strings.ToLower

	errNoPtr = errors.New("Target must be a pointer")
)

// ErrorMapper is an error which also contains a map of additional (named) errors
// which altogether caused the actual failure.
type ErrorMapper interface {
	error
	ErrorMap() map[string]error
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
		typ := reflect.TypeOf(target)
		if typ.Kind() != reflect.Ptr {
			return errNoPtr
		}

		typ = typ.Elem()
		val := reflect.ValueOf(target).Elem()
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			paramName, ok := field.Tag.Lookup(r.tag)
			if !ok {
				paramName = r.mapper(field.Name)
			}
			paramVal, ok := params[paramName]
			if !ok {
				continue
			}
			structField := val.Field(i)
			internal.ParseInto(paramVal[0], structField)
		}
	}

	return nil
}
