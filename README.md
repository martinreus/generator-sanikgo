# generator-sanikgo

An opinionated go microservice generator using the [Yeoman](https://yeoman.io) generator. Requires node to be installed.

## Before you proceed

You will need yeoman (and therefore node) + go installed. Also, [GOPATH and GOBIN](https://essential-go.programming-books.io/gopath-goroot-gobin-d6da4b8481f94757bae43be1fdfa9e73) must be correctly configured in your environment (or just add your go/bin to the PATH variable).

## What is generated?

This project has several yeoman targets, which can be called separately. However, the default target should generate a simple Go project structure.

## How to use

- `yo sanikgo`: generates a basic project with main needed files
- `yo sanikgo:[target]`: see [available targets](#available-targets)

#### Available targets

###### sanikgo:base (also invoked when not specifying any target)

Generates all base files, including `go.mod`, `cmd/main.go`, `Dockerfile` and some base folder following [go package structure](https://github.com/golang-standards/project-layout) standards.

###### sanikgo:openapi

Generates a rest API using an openapi generator (uses [oapi-codegen](https://github.com/deepmap/oapi-codegen) underneath)
