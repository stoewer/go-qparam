// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

package internal_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/stoewer/go-qparam/internal"
	"github.com/stretchr/testify/assert"
)

var now time.Time
var nowStr string

func init() {
	now = time.Now()
	b, _ := now.MarshalText()
	nowStr = string(b)
}

func TestParseInto(t *testing.T) {
	tests := []struct {
		name   string
		str    string
		target interface{}
		exp    interface{}
	}{
		{name: "test int with zero", str: "0", target: new(int), exp: int(0)},
		{name: "test int with positive val", str: "557700679194777941", target: new(int), exp: int(557700679194777941)},
		{name: "test int with negative val", str: "-403720079423501005", target: new(int), exp: int(-403720079423501005)},
		{name: "test int8 with zero", str: "0", target: new(int8), exp: int8(0)},
		{name: "test int8 with positive val", str: "75", target: new(int8), exp: int8(75)},
		{name: "test int8 with negative val", str: "-123", target: new(int8), exp: int8(-123)},
		{name: "test int16 with zero", str: "0", target: new(int16), exp: int16(0)},
		{name: "test int16 with positive val", str: "20143", target: new(int16), exp: int16(20143)},
		{name: "test int16 with negative val", str: "-31945", target: new(int16), exp: int16(-31945)},
		{name: "test int32 with zero", str: "0", target: new(int32), exp: int32(0)},
		{name: "test int32 with positive val", str: "2147481661", target: new(int32), exp: int32(2147481661)},
		{name: "test int32 with negative val", str: "-1298498081", target: new(int32), exp: int32(-1298498081)},
		{name: "test int64 with zero", str: "0", target: new(int64), exp: int64(0)},
		{name: "test int64 with positive val", str: "8674665223082153551", target: new(int64), exp: int64(8674665223082153551)},
		{name: "test int64 with negative val", str: "-6129484611666145821", target: new(int64), exp: int64(-6129484611666145821)},

		{name: "test uint with zero", str: "0", target: new(uint), exp: uint(0)},
		{name: "test uint with positive val", str: "3916589616287113937", target: new(uint), exp: uint(3916589616287113937)},
		{name: "test uint8 with zero", str: "0", target: new(uint8), exp: uint8(0)},
		{name: "test uint8 with positive val", str: "233", target: new(uint8), exp: uint8(233)},
		{name: "test uint16 with zero", str: "0", target: new(uint16), exp: uint16(0)},
		{name: "test uint16 with positive val", str: "31550", target: new(uint16), exp: uint16(31550)},
		{name: "test uint32 with zero", str: "0", target: new(uint32), exp: uint32(0)},
		{name: "test uint32 with positive val", str: "336122540", target: new(uint32), exp: uint32(336122540)},
		{name: "test uint64 with zero", str: "0", target: new(uint64), exp: uint64(0)},
		{name: "test uint64 with positive val", str: "605394647632969758", target: new(uint64), exp: uint64(605394647632969758)},

		{name: "test float32 with zero", str: "0", target: new(float32), exp: float32(0)},
		{name: "test float32 with positive val", str: "0.30091187", target: new(float32), exp: float32(0.30091187)},
		{name: "test float32 with negative val", str: "-0.51521266", target: new(float32), exp: float32(-0.51521266)},
		{name: "test float64 with zero", str: "0", target: new(float64), exp: float64(0)},
		{name: "test float64 with positive val", str: "0.8136399609900968", target: new(float64), exp: float64(0.8136399609900968)},
		{name: "test float64 with negative val", str: "-0.21426387258237492", target: new(float64), exp: float64(-0.21426387258237492)},

		{name: "test bool with true", str: "true", target: new(bool), exp: true},
		{name: "test bool with false", str: "false", target: new(bool), exp: false},

		{name: "test string", str: "foo bar", target: new(string), exp: "foo bar"},
		{name: "test time", str: nowStr, target: new(time.Time), exp: now},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			internal.ParseInto(tt.str, tt.target)
			val := reflect.ValueOf(tt.target).Elem().Interface()

			assert.Equal(t, tt.exp, val)
		})
	}
}
