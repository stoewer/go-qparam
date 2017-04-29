// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

package internal_test

import (
	"reflect"
	"testing"

	"github.com/stoewer/go-qparam/internal"
	"github.com/stoewer/go-strcase"
	"github.com/stretchr/testify/assert"
)

// TODO test returned values

type outer struct {
	FieldA    string
	FiledB    int
	C         bool    `param:"field_c"`
	Skip      int     `param:"-"`
	One       *innerA `param:"struct_one"`
	StructTwo innerB
}

type innerA struct {
	FieldD *int
	FieldE *bool
}

type innerB struct {
	FieldF      float64
	StructThree innerC
	FieldH      *float64
}

type innerC struct {
	FieldG *string
}

func TestIterator_Next(t *testing.T) {
	expected := []string{
		"field_a",
		"filed_b",
		"field_c",
		"struct_one",
		"struct_one.field_d",
		"struct_one.field_e",
		"struct_two",
		"struct_two.field_f",
		"struct_two.struct_three",
		"struct_two.struct_three.field_g",
		"struct_two.field_h",
	}

	data := &outer{}
	it := internal.NewIterator(reflect.ValueOf(data), "param", strcase.SnakeCase)

	index := 0
	for i := 0; i < len(expected); i++ {
		name, _ := it.Next()
		assert.Equal(t, expected[i], name)
		index++
	}
}

func TestIterator_HasNext(t *testing.T) {
	expected := []string{
		"field_a",
		"filed_b",
		"field_c",
		"struct_one",
		"struct_one.field_d",
		"struct_one.field_e",
		"struct_two",
		"struct_two.field_f",
		"struct_two.struct_three",
		"struct_two.struct_three.field_g",
		"struct_two.field_h",
	}

	data := &outer{}
	it := internal.NewIterator(reflect.ValueOf(data), "param", strcase.SnakeCase)

	i := 0
	for it.HasNext() {
		name, _ := it.Next()
		assert.Equal(t, expected[i], name)
		i++
	}
}

func TestIterator_Skip(t *testing.T) {
	data := &outer{}
	it := internal.NewIterator(reflect.ValueOf(data), "param", strcase.SnakeCase)

	var name string
	for i := 0; i < 4; i++ {
		name, _ = it.Next()
	}

	assert.Equal(t, "struct_one", name)
	it.SkipStruct()

	for i := 0; i < 3; i++ {
		name, _ = it.Next()
	}

	assert.Equal(t, "struct_two.struct_three", name)
	it.SkipStruct()
	it.SkipStruct()

	name, _ = it.Next()
	assert.Equal(t, "struct_two.field_h", name)
}
