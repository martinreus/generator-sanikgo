# generator-sanikgo

An opinionated go microservice generator using the [Yeoman](https://yeoman.io) generator. Requires node to be installed.

## What is generated?

This project has several yeoman targets, which can be called separately. However, the default target should generate a complete go microservice.

#### Available targets

###### sanikgo:base

Generates all base files, including `go.mod`, `cmd/main.go`, `Dockerfile` and some base folder following [go package structure](https://github.com/golang-standards/project-layout) standards.

###### sanikgo:openapi

Generates a rest API using an openapi generator (uses [oapi-codegen](https://github.com/deepmap/oapi-codegen) underneath)
