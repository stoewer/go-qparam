[![Build Status](https://travis-ci.org/stoewer/go-qparam.svg?branch=master)](https://travis-ci.org/stoewer/go-qparam)
[![Coverage Status](https://coveralls.io/repos/github/stoewer/go-qparam/badge.svg?branch=master)](https://coveralls.io/github/stoewer/go-qparam?branch=master)
[![GoDoc](https://godoc.org/github.com/stoewer/go-qparam?status.svg)](https://godoc.org/github.com/stoewer/go-qparam)
---

# Go qparam

Package qparam provides convenient functions to read query parameters or form values into the fields
of one (or more) target struct.

The following field types are supported: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, 
`uint16`, `uint32`, `uint64`, `float32`, `float64`, `bool`, `string`. In addition the package handles 
also all types implementing the `TextUnmarshaler` interface from the `encoding` package. Furthermore 
pointers and slices of all before mentioned types are supported.

To handle hierarchically structured data, the package can also be used to read values into fields
of nested structs. In such a case the keys of the source must use dots (`.`) as 'path' delimiter. 
The reader can further be configured to use custom field tags and a custom name mapping, which keeps 
the necessity to add tags to struct fields at a minimum.

## Versions and stability

The package is not stable yet.

## Install and use

```sh
go get -u github.com/stoewer/go-qparam
```

The following example show how `qparam` can be used (there are more examples in the [reference](https://godoc.org/github.com/stoewer/go-qparam)):

```go
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
```
