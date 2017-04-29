// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

package qparam_test

import (
	"net/url"
	"testing"
	"time"

	"github.com/stoewer/go-qparam"
	"github.com/stoewer/go-strcase"
	"github.com/stretchr/testify/assert"
)

var (
	now    time.Time
	nowStr string
)

func init() {
	b, _ := time.Now().MarshalText()
	nowStr = string(b)
	now.UnmarshalText(b)
}

type noTag struct {
	Int    int8
	Uint   uint
	Float  float64
	String string
	Inner  withTag
	Bool   bool
	Time   time.Time
}

type withTag struct {
	Int    int        `param:"a"`
	Uint   uint       `param:"b"`
	Float  float64    `param:"c"`
	String string     `param:"d"`
	Time   *time.Time `param:"f"`
	Bool   bool       `param:"e"`
}

func TestReader_Read(t *testing.T) {

	type pointers struct {
		Int32Ptr  *int32
		Uint32Ptr *uint32
	}
	type times struct {
		Time    time.Time
		TimePtr *time.Time
	}

	type strings struct {
		String    string
		StringPtr *string
	}

	type test struct {
		Int      int
		Int8     int8
		Int16    int16
		Int32    int32
		Int64    int64
		Uint     uint
		Uint8    uint8
		Uint16   uint16
		Uint32   uint32
		Uint64   uint64
		Pointers *pointers
		Bool     bool
		Times    times
		Strings  *strings
	}

	expected := test{
		Int:    -345,
		Int8:   33,
		Int16:  -10344,
		Int32:  32999,
		Int64:  -939393,
		Uint:   10231,
		Uint8:  177,
		Uint16: 533,
		Uint32: 99,
		Uint64: 192837,
		Pointers: &pointers{
			Int32Ptr:  new(int32),
			Uint32Ptr: new(uint32),
		},
		Bool: true,
		Times: times{
			Time:    now,
			TimePtr: &now,
		},
		Strings: &strings{
			String:    "foo",
			StringPtr: new(string),
		},
	}
	*expected.Pointers.Int32Ptr = -253
	*expected.Pointers.Uint32Ptr = 94883
	*expected.Strings.StringPtr = "bar"

	target := test{}
	values := url.Values{
		"int":                []string{"-345"},
		"int8":               []string{"33"},
		"int16":              []string{"-10344"},
		"int32":              []string{"32999"},
		"int64":              []string{"-939393"},
		"uint":               []string{"10231"},
		"uint8":              []string{"177"},
		"uint16":             []string{"533"},
		"uint32":             []string{"99"},
		"uint64":             []string{"192837"},
		"pointers.int32ptr":  []string{"-253"},
		"pointers.uint32ptr": []string{"94883"},
		"bool":               []string{"true"},
		"times.time":         []string{nowStr},
		"times.time_ptr":     []string{nowStr},
		"strings.string":     []string{"foo"},
		"strings.string_ptr": []string{"bar"},
	}

	reader := qparam.New(qparam.Mapper(strcase.SnakeCase), qparam.Tag("param"))
	err := reader.Read(values, &target)

	assert.NoError(t, err)
	assert.EqualValues(t, expected, target)
}
