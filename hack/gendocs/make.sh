#!/usr/bin/env bash

pushd $GOPATH/src/github.com/appscode/analytics/hack/gendocs
go run main.go
popd
