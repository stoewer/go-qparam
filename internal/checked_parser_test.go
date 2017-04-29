// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

package internal

import (
	"reflect"
	"testing"
	"time"

	"github.com/guregu/null"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	now      time.Time
	nowStr   string
	nullTrue = null.BoolFrom(true)
)

func init() {
	b, _ := time.Now().MarshalText()
	nowStr = string(b)
	now.UnmarshalText(b)
}

func TestSelectCheckedParser(t *testing.T) {
	data := []struct {
		Value       string
		Target      interface{}
		Expected    interface{}
		ExpectedErr bool
	}{
		{Value: nowStr, Target: &time.Time{}, Expected: &now},
		{Value: "not a time", Target: &time.Time{}, ExpectedErr: true},

		{Value: "true", Target: &null.Bool{}, Expected: &nullTrue},
		{Value: "not a bool", Target: &null.Bool{}, ExpectedErr: true},
	}

	for _, tt := range data {
		parser, ok := SelectCheckedParser(reflect.ValueOf(tt.Target))
		require.True(t, ok, "no parser found")

		err := parser.Parse(reflect.ValueOf(tt.Target), tt.Value)
		if tt.ExpectedErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tt.Expected, tt.Target)
		}
	}
}
