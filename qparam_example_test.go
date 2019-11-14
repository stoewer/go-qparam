// Copyright (c) 2017, A. Stoewer <adrian@stoewer.me>
// All rights reserved.

package qparam_test

import (
	"fmt"
	"net/url"

	"github.com/stoewer/go-qparam"
	"github.com/stoewer/go-strcase"
)

func Example_simple() {
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

	reader := qparam.NewReader()
	_ = reader.Read(values, &page, &filters)

	fmt.Println(page.Limit, page.Offset, filters.Name, filters.Age)
	// Output: 25 100 Doe 31
}

func Example_nested() {
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

	reader := qparam.NewReader()
	_ = reader.Read(values, &contact)

	fmt.Println(contact.Name, contact.Phone.Label, contact.Phone.Number)
	// Output: John, Doe Mobile +33 112 33445566
}

func Example_tag() {
	values := url.Values{"session_id": []string{"abcdefghijklmn"}}
	info := struct {
		Session string `mytag:"session_id"`
	}{}

	reader := qparam.NewReader(qparam.Tag("mytag"))
	_ = reader.Read(values, &info)

	fmt.Println(info.Session)
	// Output: abcdefghijklmn
}

func Example_mapper() {
	values := url.Values{"session_id": []string{"abcdefghijklmn"}}
	info := struct {
		SessionID string
	}{}

	reader := qparam.NewReader(qparam.Mapper(strcase.SnakeCase))
	_ = reader.Read(values, &info)

	fmt.Println(info.SessionID)
	// Output: abcdefghijklmn
}
