#!/bin/bash

set -ue

pushd $(dirname $0)
mkdir -p bin
pushd bin

if ! wget https://github.com/Yandex-Practicum/go-autotests/releases/download/v0.11.2/shortenertest
then
	echo "failed to get shortenertest binary"
	exit 1
fi


chmod 755 shortenertest

popd
popd
