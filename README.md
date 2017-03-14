# Go qparam

The package `qparam` provides convenient functions to read query parameters from request 
URLs. On errors the package functions create errors which provide detailed information about
missing parameters and other failure causes.

## Versioning and stability

Although the master branch is supposed to remain always backward compatible, the repository
contains version tags in order to support vendoring tools such as `glide`.
The tag names follow semantic versioning conventions and have the following format `v1.0.0`.

## Install and use

```sh
go get -u github.com/stoewer/go-qparam
```
