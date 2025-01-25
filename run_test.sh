#!/bin/bash

set -ue

ITER="${1:?}"
ITER_ARG="^TestIteration${ITER}$"

bash build.sh

./bin/shortenertest \
	-test.v \
	-test.run="$ITER_ARG" \
	-binary-path=cmd/shortener/shortener
