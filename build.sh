#!/usr/bin/env bash

VERSION_CODE=`git log | head -n 1 | awk '{print $2}'`

go build -ldflags "-X 'github.com/Moekr/lightning/util/version.Code=$VERSION_CODE'" \
    -a -v -o output/lightning