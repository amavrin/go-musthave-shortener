#!/bin/bash

set -ue

DIR=`dirname $0`
pushd "${DIR}/cmd/shortener"
go build -o shortener *.go
popd
