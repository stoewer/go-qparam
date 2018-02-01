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
	"github.com/stretchr/testify/require"
)

var (
	now          time.Time
	nowStr       string
	yesterday    time.Time
	yesterdayStr string
	tomorrow     time.Time
	tomorrowStr  string
)

func init() {
	b, _ := time.Now().MarshalText()
	nowStr = string(b)
	now.UnmarshalText(b)

	b, _ = time.Now().Add(-24 * time.Hour).MarshalText()
	yesterdayStr = string(b)
	yesterday.UnmarshalText(b)

	b, _ = time.Now().Add(24 * time.Hour).MarshalText()
	tomorrowStr = string(b)
	tomorrow.UnmarshalText(b)
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

	type slices struct {
		IntSlice     []int
		IntPtrSlice  []*int
		TimeSlice    []time.Time
		TimePtrSlice []*time.Time
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

	t.Run("no pointer error", func(t *testing.T) {
		target := times{}

		reader := qparam.NewReader()
		err := reader.Read(url.Values{}, target)

		assert.Error(t, err)
	})

	t.Run("no struct error", func(t *testing.T) {
		target := "not a struct"

		reader := qparam.NewReader()
		err := reader.Read(url.Values{}, &target)

		assert.Error(t, err)
	})

	t.Run("unsupported type", func(t *testing.T) {
		target := struct {
			Field map[string]string
			Slice []map[string]string
		}{Field: make(map[string]string), Slice: make([]map[string]string, 0)}
		values := url.Values{"field": []string{"map"}, "slice": []string{"map"}}

		reader := qparam.NewReader()
		err := reader.Read(values, &target)

		assert.Error(t, err)
	})

	t.Run("parsing errors", func(t *testing.T) {
		timesTarget := times{}
		slicesTarget := slices{}
		pointersTarget := pointers{}
		values := url.Values{
			"time": []string{"not a time"}, "time_ptr": []string{nowStr, tomorrowStr},
			"int_slice": []string{"not an int", "2", "3"}, "time_slice": []string{"not a time"},
			"int32ptr": []string{"not and int"}, "uint32ptr": []string{"-399"},
		}

		reader := qparam.NewReader(qparam.Mapper(strcase.SnakeCase), qparam.Tag("param"))
		err := reader.Read(values, &timesTarget, &slicesTarget, &pointersTarget)

		assert.EqualError(t, err, "6 errors occurred while reading parameters")
		multi, ok := err.(qparam.MultiError)
		require.True(t, ok, "not a MultiError")
		assert.Equal(t, 6, len(multi.ErrorMap()))
	})

	t.Run("struct with slices", func(t *testing.T) {
		one := 1
		two := 2
		slicesExpected := slices{
			IntSlice:     []int{1, 2, 3},
			IntPtrSlice:  []*int{&one, &two, &one, &two},
			TimeSlice:    []time.Time{yesterday, now, tomorrow},
			TimePtrSlice: []*time.Time{&yesterday, &now, &tomorrow},
		}
		values := url.Values{
			"int_slice":      []string{"1", "2", "3"},
			"int_ptr_slice":  []string{"1", "2", "1", "2"},
			"time_slice":     []string{yesterdayStr, nowStr, tomorrowStr},
			"time_ptr_slice": []string{yesterdayStr, nowStr, tomorrowStr},
		}

		slicesTarget := slices{}
		reader := qparam.NewReader(qparam.Mapper(strcase.SnakeCase), qparam.Tag("param"))
		err := reader.Read(values, &slicesTarget)

		assert.NoError(t, err)
		assert.EqualValues(t, slicesExpected, slicesTarget)
	})

	t.Run("strict mode", func(t *testing.T) {
		values := url.Values{
			"time":     []string{nowStr},
			"time_ptr": []string{nowStr},
			"foo":      []string{"not expected"},
		}

		timesTarget := times{}
		reader := qparam.NewReader(qparam.Mapper(strcase.SnakeCase), qparam.Strict(true))
		err := reader.Read(values, &timesTarget)

		assert.Error(t, err)
		multi, ok := err.(qparam.MultiError)
		require.True(t, ok, "not a MultiError")
		_, ok = multi.ErrorMap()["foo"]
		assert.True(t, ok, "field foo not in error")
	})

	t.Run("multiple structs", func(t *testing.T) {
		str := "foo"
		timesExpected := times{Time: now, TimePtr: &now}
		stringExpected := strings{String: "bar", StringPtr: &str}
		values := url.Values{
			"time":       []string{nowStr},
			"time_ptr":   []string{nowStr},
			"string":     []string{"bar"},
			"string_ptr": []string{"foo"},
		}

		timesTarget := times{}
		stringsTarget := strings{}
		reader := qparam.NewReader(qparam.Mapper(strcase.SnakeCase), qparam.Tag("param"))
		err := reader.Read(values, &timesTarget, &stringsTarget)

		assert.NoError(t, err)
		assert.EqualValues(t, timesExpected, timesTarget)
		assert.EqualValues(t, stringExpected, stringsTarget)
	})

	t.Run("nested structs", func(t *testing.T) {
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
		target := test{Pointers: &pointers{}, Strings: &strings{}}
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

		reader := qparam.NewReader(qparam.Mapper(strcase.SnakeCase), qparam.Tag("param"))
		err := reader.Read(values, &target)

		assert.NoError(t, err)
		assert.EqualValues(t, expected, target)
	})
}
