[![CircleCI](https://circleci.com/gh/stoewer/go-qparam/tree/master.svg?style=svg)](https://circleci.com/gh/stoewer/go-qparam/tree/master)
[![codecov](https://codecov.io/gh/stoewer/go-qparam/branch/master/graph/badge.svg)](https://codecov.io/gh/stoewer/go-qparam)
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

This package can be considered stable and ready to use. All releases follow the rules of
[semantic versioning](http://semver.org).

Although the master branch is supposed to remain stable, there is not guarantee that braking changes will not
be merged into master when major versions are released. Therefore the repository contains version tags in
order to support vendoring tools such as [glide](https://glide.sh). The tag names follow common conventions
and have the following format `v1.0.0`. This package also supports Go modules introduced with version 1.11.

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
reader := qparam.NewReader()
reader.Read(values, &page, &filters)
```
