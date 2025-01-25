#!/bin/bash

set -ue

pushd `dirname $0`
mkdir -p bin
pushd bin

wget https://github.com/Yandex-Practicum/go-autotests/releases/download/v0.11.2/shortenertest
chmod 755 shortenertest

popd
popd
