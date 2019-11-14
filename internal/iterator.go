// Copyright (c) 2017, A. Stoewer <adrian@stoewer.me>
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

type state byte

const (
	forwardRequired state = iota
	ok
	done
)

// NewIterator returns a new Iterator. The returned iterator can be used exactly one time to iterate over
// fields of the target struct.
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

// Iterator is used to iterate over struct fields and the fields of its child structs.
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

// HasNext indicates whether or not the iterator can return an additional field. HasNext should always be
// called prior to calling the Next method.
func (it *Iterator) HasNext() bool {
	if it.state == forwardRequired {
		it.state = it.forward()
	}
	return it.state != done
}

// Next returns the next field (along with its path/name). The method HasNext should be used in order to
// check whether it is safe to call Next or not.
func (it *Iterator) Next() (string, reflect.Value) {
	if it.state == forwardRequired {
		it.state = it.forward()
	}
	if it.state == ok {
		it.state = forwardRequired
	}
	return it.fieldPath, it.fieldValue
}

// SkipStruct skips all remaining fields of the current struct. This will end the iteration or will continue with
// the next field of the parent struct.
func (it *Iterator) SkipStruct() {
	it.state = forwardRequired
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

			return done
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

		if it.fieldValue.Kind() == reflect.Ptr && !it.fieldValue.IsNil() {
			it.fieldValue = it.fieldValue.Elem()
		}

		// if the field is a struct: go to next level
		if it.fieldValue.Kind() == reflect.Struct {
			it.parents = append(it.parents, parent{current: it.current, index: it.index, name: fieldName})
			it.current = it.fieldValue
			it.index = 0
			return ok
		}

		// forwarding complete
		it.index++
		return ok
	}
}
