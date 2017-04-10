// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

package qparam_test

import (
	"net/url"
	"testing"

	"github.com/stoewer/go-qparam"
	"github.com/stretchr/testify/assert"
)

func TestReader_Read(t *testing.T) {
	type strings struct {
		Foo string
		Bar string
	}

	tests := []struct {
		Values   url.Values
		Expected strings
	}{
		{Values: url.Values{"foo": []string{"FOO"}, "bar": []string{"BAR"}}, Expected: strings{Foo: "FOO", Bar: "BAR"}},
	}

	for _, tt := range tests {
		r := qparam.New()
		target := strings{}
		err := r.Read(tt.Values, &target)
		assert.NoError(t, err)
		assert.Equal(t, target, tt.Expected)
	}
}
