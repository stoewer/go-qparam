// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

package internal

import (
	"reflect"
	"strings"
)

type parent struct {
	current reflect.Value
	index   int
	name    string
}

type state int

const (
	FORWARD_REQUIRED state = iota
	OK
	DONE
)

func NewIterator(target reflect.Value, tag string, mapper func(string) string) *Iterator {
	if target.Kind() == reflect.Ptr {
		target = target.Elem()
	}

	return &Iterator{
		tag:     tag,
		mapper:  mapper,
		current: target,
	}
}

type Iterator struct {
	tag        string
	mapper     func(string) string
	current    reflect.Value
	parents    []parent
	index      int
	state      state
	fieldValue reflect.Value
	fieldPath  string
}

func (it *Iterator) HasNext() bool {
	if it.state == FORWARD_REQUIRED {
		it.state = it.forward()
	}
	return it.state != DONE
}

func (it *Iterator) Next() (string, reflect.Value) {
	if it.state == FORWARD_REQUIRED {
		it.state = it.forward()
	}
	if it.state == OK {
		it.state = FORWARD_REQUIRED
	}
	return it.fieldPath, it.fieldValue
}

func (it *Iterator) SkipStruct() {
	it.state = FORWARD_REQUIRED
	it.index = it.current.NumField()
}

func (it *Iterator) forward() state {
	for {
		// check end condition
		if it.index >= it.current.NumField() {
			if len(it.parents) > 0 {
				parent := it.parents[len(it.parents)-1]
				it.current = parent.current
				it.index = parent.index + 1
				it.parents = it.parents[:len(it.parents)-1]
				continue
			}

			return DONE
		}

		// skip if not settable
		it.fieldValue = it.current.Field(it.index)
		if !it.fieldValue.CanSet() {
			it.index++
			continue
		}

		// skip if tag is "-"
		field := it.current.Type().Field(it.index)
		fieldName := field.Tag.Get(it.tag)
		if fieldName == "-" {
			it.index++
			continue
		}

		if fieldName == "" {
			fieldName = it.mapper(field.Name)
		}

		// create empty field elements
		if it.fieldValue.Kind() == reflect.Ptr {
			if it.fieldValue.IsNil() {
				it.fieldValue.Set(reflect.New(it.fieldValue.Type().Elem()))
			}
			it.fieldValue = it.fieldValue.Elem()
		}

		// determine field path
		if len(it.parents) > 0 {
			parentNames := make([]string, 0, len(it.parents)+1)
			for _, item := range it.parents {
				parentNames = append(parentNames, item.name)
			}
			it.fieldPath = strings.Join(append(parentNames, fieldName), ".")
		} else {
			it.fieldPath = fieldName
		}

		// if the field is a struct: go to next level
		if it.fieldValue.Kind() == reflect.Struct {
			it.parents = append(it.parents, parent{current: it.current, index: it.index, name: fieldName})
			it.current = it.fieldValue
			it.index = 0
			return OK
		}

		// forwarding complete
		it.index++
		return OK
	}
}
