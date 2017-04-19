package internal

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntParser_Parse(t *testing.T) {
	data := []struct {
		Name        string
		Value       string
		Expected    reflect.Value
		ExpectedErr error
	}{
		{Name: "zero", Value: "0", Expected: reflect.ValueOf(0)},
		{Name: "positive int", Value: "8674665223082153551", Expected: reflect.ValueOf(8674665223082153551)},
		{Name: "negative int", Value: "-6129484611666145821", Expected: reflect.ValueOf(-6129484611666145821)},
		{Name: "invalid", Value: "invalid", ExpectedErr: errors.New(`strconv.ParseInt: parsing "invalid": invalid syntax`)},
	}

	for _, tt := range data {
		t.Run(tt.Name, func(t *testing.T) {
			val, err := IntParser.Parse(tt.Value)
			if tt.ExpectedErr != nil {
				assert.EqualError(t, err, tt.ExpectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.Expected.Int(), val.Int())
			}
		})
	}
}

func TestUintParser_Parse(t *testing.T) {
	data := []struct {
		Name        string
		Value       string
		Expected    reflect.Value
		ExpectedErr error
	}{
		{Name: "zero", Value: "0", Expected: reflect.ValueOf(uint(0))},
		{Name: "positive int", Value: "8674665223082153551", Expected: reflect.ValueOf(uint(8674665223082153551))},
		{Name: "invalid", Value: "invalid", ExpectedErr: errors.New(`strconv.ParseUint: parsing "invalid": invalid syntax`)},
	}

	for _, tt := range data {
		t.Run(tt.Name, func(t *testing.T) {
			val, err := UintParser.Parse(tt.Value)
			if tt.ExpectedErr != nil {
				assert.EqualError(t, err, tt.ExpectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.Expected.Uint(), val.Uint())
			}
		})
	}
}

func TestFloatParser_Parse(t *testing.T) {
	data := []struct {
		Name        string
		Value       string
		Expected    reflect.Value
		ExpectedErr error
	}{
		{Name: "zero", Value: "0", Expected: reflect.ValueOf(0.0)},
		{Name: "positive float", Value: "0.8136399609900968", Expected: reflect.ValueOf(0.8136399609900968)},
		{Name: "negative float", Value: "-0.21426387258237492", Expected: reflect.ValueOf(-0.21426387258237492)},
		{Name: "positive int", Value: "2147481661", Expected: reflect.ValueOf(float64(2147481661))},
		{Name: "negative int", Value: "-1298498081", Expected: reflect.ValueOf(float64(-1298498081))},
		{Name: "invalid", Value: "invalid", ExpectedErr: errors.New(`strconv.ParseFloat: parsing "invalid": invalid syntax`)},
	}

	for _, tt := range data {
		t.Run(tt.Name, func(t *testing.T) {
			val, err := FloatParser.Parse(tt.Value)
			if tt.ExpectedErr != nil {
				assert.EqualError(t, err, tt.ExpectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.Expected.Float(), val.Float())
			}
		})
	}
}

func TestBoolParser_Parse(t *testing.T) {
	data := []struct {
		Name        string
		Value       string
		Expected    reflect.Value
		ExpectedErr error
	}{
		{Name: "true", Value: "true", Expected: reflect.ValueOf(true)},
		{Name: "false", Value: "false", Expected: reflect.ValueOf(false)},
		{Name: "invalid", Value: "invalid", ExpectedErr: errors.New(`strconv.ParseBool: parsing "invalid": invalid syntax`)},
	}

	for _, tt := range data {
		t.Run(tt.Name, func(t *testing.T) {
			val, err := BoolParser.Parse(tt.Value)
			if tt.ExpectedErr != nil {
				assert.EqualError(t, err, tt.ExpectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.Expected.Bool(), val.Bool())
			}
		})
	}
}

func TestStringParser_Parse(t *testing.T) {
	data := []struct {
		Name        string
		Value       string
		Expected    reflect.Value
		ExpectedErr error
	}{
		{Name: "foo", Value: "foo", Expected: reflect.ValueOf("foo")},
		{Name: "bar", Value: "bar", Expected: reflect.ValueOf("bar")},
	}

	for _, tt := range data {
		t.Run(tt.Name, func(t *testing.T) {
			val, err := StringParser.Parse(tt.Value)
			assert.NoError(t, err)
			assert.Equal(t, tt.Expected.String(), val.String())
		})
	}
}
