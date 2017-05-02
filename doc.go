// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

/*
Package qparam provides convenient functions to read query parameters or form values into the fields
of one (or more) target struct:

	type Page struct {
		Limit  int
		Offset int
	}

	type Filters struct {
		Name string
		Age  int
	}

	values := map[string][]string{
		"limit":  {"25"},
		"offset": {"100"},
		"name":   {"Doe"},
		"age":    {"31"},
	}

	var page Page
	var filters Filters
	reader := qparam.New()
	reader.Read(values, &page, &filters)


The following field types are supported: int, int8, int16, int32, int64, uint, uint8, uint16,
uint32, uint64, float32, float64, bool, string. In addition the package handles also all types
implementing the TextUnmarshaler interface from the encoding package. Furthermore pointers and
slices of all before mentioned types are supported.

To handle hierarchically structured data, the package can also be used to read values into fields
of nested structs. In such a case the keys of the source must use dots as some kind of path
delimiter:

	type Phone struct {
		Label  string
		Number string
	}

	type Contact struct {
		Name  string
		Phone Phone
	}

	values := map[string][]string{
		"name":         {"John, Doe"},
		"phone.label":  {"Mobile"},
		"phone.number": {"+33 112 33445566"},
	}

	var contact Contact
	reader := qparam.New()
	reader.Read(values, &contact)

The reader can further be configured to use custom field tags and a custom name mapping, which keeps
the necessity to add tags to struct fields at a minimum (check the examples for more details).
*/
package qparam
