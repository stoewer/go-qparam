// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

package internal_test

import (
	"reflect"
	"testing"

	"github.com/stoewer/go-qparam/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSelectParser_Int(t *testing.T) {
	data := []struct {
		Name        string
		Value       string
		Expected    reflect.Value
		ExpectedErr bool
	}{
		{Name: "zero int", Value: "0", Expected: reflect.ValueOf(int(0))},
		{Name: "positive int", Value: "8674665223082153551", Expected: reflect.ValueOf(int(8674665223082153551))},
		{Name: "negative int", Value: "-6129484611666145821", Expected: reflect.ValueOf(int(-6129484611666145821))},
		{Name: "invalid int", Value: "invalid", Expected: reflect.ValueOf(int(0)), ExpectedErr: true},

		{Name: "zero int8", Value: "0", Expected: reflect.ValueOf(int8(0))},
		{Name: "positive int8", Value: "122", Expected: reflect.ValueOf(int8(122))},
		{Name: "negative int8", Value: "-84", Expected: reflect.ValueOf(int8(-84))},
		{Name: "invalid int8", Value: "invalid", Expected: reflect.ValueOf(int8(0)), ExpectedErr: true},

		{Name: "zero int16", Value: "0", Expected: reflect.ValueOf(int16(0))},
		{Name: "positive int16", Value: "24937", Expected: reflect.ValueOf(int16(24937))},
		{Name: "negative int16", Value: "-19253", Expected: reflect.ValueOf(int16(-19253))},
		{Name: "invalid int16", Value: "invalid", Expected: reflect.ValueOf(int16(0)), ExpectedErr: true},

		{Name: "zero int32", Value: "0", Expected: reflect.ValueOf(int32(0))},
		{Name: "positive int32", Value: "1427131847", Expected: reflect.ValueOf(int32(1427131847))},
		{Name: "negative int32", Value: "-939984059", Expected: reflect.ValueOf(int32(-939984059))},
		{Name: "invalid int32", Value: "invalid", Expected: reflect.ValueOf(int32(0)), ExpectedErr: true},

		{Name: "zero int64", Value: "0", Expected: reflect.ValueOf(int64(0))},
		{Name: "positive int64", Value: "3916589616287113937", Expected: reflect.ValueOf(int64(3916589616287113937))},
		{Name: "negative int64", Value: "-6334824724549167320", Expected: reflect.ValueOf(int64(-6334824724549167320))},
		{Name: "invalid int64", Value: "invalid", Expected: reflect.ValueOf(int64(0)), ExpectedErr: true},
	}

	for _, tt := range data {
		t.Run(tt.Name, func(t *testing.T) {
			parser, ok := internal.SelectParser(tt.Expected)
			require.True(t, ok)

			val, err := parser.Parse(tt.Value)
			if tt.ExpectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.Expected.Int(), val.Int())
			}
		})
	}
}

func TestSelectParser_Uint(t *testing.T) {
	data := []struct {
		Name        string
		Value       string
		Expected    reflect.Value
		ExpectedErr bool
	}{
		{Name: "zero uint", Value: "0", Expected: reflect.ValueOf(uint(0))},
		{Name: "positive uint", Value: "646203300", Expected: reflect.ValueOf(uint(646203300))},
		{Name: "invalid uint", Value: "invalid", Expected: reflect.ValueOf(uint(0)), ExpectedErr: true},

		{Name: "zero uint8", Value: "0", Expected: reflect.ValueOf(uint8(0))},
		{Name: "positive uint8", Value: "221", Expected: reflect.ValueOf(uint8(221))},
		{Name: "invalid uint8", Value: "invalid", Expected: reflect.ValueOf(uint8(0)), ExpectedErr: true},

		{Name: "zero uint16", Value: "0", Expected: reflect.ValueOf(uint16(0))},
		{Name: "positive uint16", Value: "24537", Expected: reflect.ValueOf(uint16(24537))},
		{Name: "invalid uint16", Value: "invalid", Expected: reflect.ValueOf(uint16(0)), ExpectedErr: true},

		{Name: "zero uint32", Value: "0", Expected: reflect.ValueOf(uint32(0))},
		{Name: "positive uint32", Value: "1106410694", Expected: reflect.ValueOf(uint32(1106410694))},
		{Name: "invalid uint32", Value: "invalid", Expected: reflect.ValueOf(uint32(0)), ExpectedErr: true},

		{Name: "zero uint64", Value: "0", Expected: reflect.ValueOf(uint64(0))},
		{Name: "positive uint64", Value: "894385949183117216", Expected: reflect.ValueOf(uint64(894385949183117216))},
		{Name: "invalid uint64", Value: "invalid", Expected: reflect.ValueOf(uint64(0)), ExpectedErr: true},
	}

	for _, tt := range data {
		t.Run(tt.Name, func(t *testing.T) {
			parser, ok := internal.SelectParser(tt.Expected)
			require.True(t, ok)

			val, err := parser.Parse(tt.Value)
			if tt.ExpectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.Expected.Uint(), val.Uint())
			}
		})
	}
}

func TestSelectParser_Float(t *testing.T) {
	data := []struct {
		Name        string
		Value       string
		Expected    reflect.Value
		ExpectedErr bool
	}{
		{Name: "zero float32", Value: "0", Expected: reflect.ValueOf(float32(0.0))},
		{Name: "positive float32", Value: "0.81363994", Expected: reflect.ValueOf(float32(0.81363994))},
		{Name: "negative float32", Value: "-0.21426387", Expected: reflect.ValueOf(float32(-0.21426387))},
		{Name: "invalid float32", Value: "invalid", Expected: reflect.ValueOf(float32(0.0)), ExpectedErr: true},

		{Name: "zero float64", Value: "0", Expected: reflect.ValueOf(float64(0.0))},
		{Name: "positive float64", Value: "0.4688898449024232", Expected: reflect.ValueOf(float64(0.4688898449024232))},
		{Name: "negative float64", Value: "-0.28303415118044517", Expected: reflect.ValueOf(float64(-0.28303415118044517))},
		{Name: "invalid float64", Value: "invalid", Expected: reflect.ValueOf(float64(0.0)), ExpectedErr: true},
	}

	for _, tt := range data {
		t.Run(tt.Name, func(t *testing.T) {
			parser, ok := internal.SelectParser(tt.Expected)
			require.True(t, ok)

			val, err := parser.Parse(tt.Value)
			if tt.ExpectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.Expected.Float(), val.Float())
			}
		})
	}
}

func TestSelectParser_Bool(t *testing.T) {
	data := []struct {
		Name        string
		Value       string
		Expected    reflect.Value
		ExpectedErr bool
	}{
		{Name: "true", Value: "true", Expected: reflect.ValueOf(true)},
		{Name: "false", Value: "false", Expected: reflect.ValueOf(false)},
		{Name: "invalid", Value: "invalid", Expected: reflect.ValueOf(true), ExpectedErr: true},
	}

	for _, tt := range data {
		t.Run(tt.Name, func(t *testing.T) {
			parser, ok := internal.SelectParser(tt.Expected)
			require.True(t, ok)

			val, err := parser.Parse(tt.Value)
			if tt.ExpectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.Expected.Bool(), val.Bool())
			}
		})
	}
}

func TestSelectParser_String(t *testing.T) {
	data := []struct {
		Name        string
		Value       string
		Expected    reflect.Value
		ExpectedErr bool
	}{
		{Name: "foo", Value: "foo", Expected: reflect.ValueOf("foo")},
		{Name: "bar", Value: "bar", Expected: reflect.ValueOf("bar")},
	}

	for _, tt := range data {
		t.Run(tt.Name, func(t *testing.T) {
			parser, ok := internal.SelectParser(tt.Expected)
			require.True(t, ok)

			val, err := parser.Parse(tt.Value)
			if tt.ExpectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.Expected.String(), val.String())
			}
		})
	}
}
