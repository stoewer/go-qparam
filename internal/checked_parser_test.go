package internal

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/guregu/null"
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

func TestTextParser_Check(t *testing.T) {
	data := []struct {
		Value    reflect.Value
		Expected bool
	}{
		{Value: reflect.ValueOf(&now), Expected: true},
		{Value: reflect.ValueOf(&null.String{}), Expected: true},
		{Value: reflect.ValueOf(&null.Bool{}), Expected: true},
		{Value: reflect.ValueOf(&sql.NullString{}), Expected: false},
	}

	textParser := &TextParser{}
	for _, tt := range data {
		assert.Equal(t, tt.Expected, textParser.Check(tt.Value))
	}
}

func TestTextParser_Parse(t *testing.T) {
	data := []struct {
		Value       string
		Target      interface{}
		Expected    interface{}
		ExpectedErr error
	}{
		{Value: nowStr, Target: &time.Time{}, Expected: &now},
		{Value: "not a time", Target: &time.Time{}, ExpectedErr: errors.New("Error while calling UnmarshalText")},
	}

	textParser := &TextParser{}
	for _, tt := range data {
		err := textParser.Parse(reflect.ValueOf(tt.Target), tt.Value)
		if tt.ExpectedErr != nil {
			assert.EqualError(t, err, tt.ExpectedErr.Error())
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tt.Expected, tt.Target)
		}
	}
}
