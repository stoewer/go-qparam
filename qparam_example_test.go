// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

package qparam_test

import (
	"fmt"
	"net/url"

	"github.com/stoewer/go-qparam"
	"github.com/stoewer/go-strcase"
)

func ExampleSimple() {
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

	fmt.Printf("limit: %d, offset: %d, name: %s, age: %d\n", page.Limit, page.Offset, filters.Name, filters.Age)
	// Output: limit: 25, offset: 100, name: Doe, age: 31
}

func ExampleNestedStructs() {
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

	fmt.Printf("name: %s, label: %s, number: %s", contact.Name, contact.Phone.Label, contact.Phone.Number)
	// Output: name: John, Doe, label: Mobile, number: +33 112 33445566
}

func ExampleCustomTag() {
	values := url.Values{"session_id": []string{"abcdefghijklmn"}}
	info := struct {
		Session string `mytag:"session_id"`
	}{}

	reader := qparam.New(qparam.Tag("mytag"))
	reader.Read(values, &info)

	fmt.Printf("session: %s\n", info.Session)
	// Output: session: abcdefghijklmn
}

func ExampleCustomMapper() {
	values := url.Values{"session_id": []string{"abcdefghijklmn"}}
	info := struct {
		SessionID string
	}{}

	reader := qparam.New(qparam.Mapper(strcase.SnakeCase))
	reader.Read(values, &info)

	fmt.Printf("session: %s\n", info.SessionID)
	// Output: session: abcdefghijklmn
}
