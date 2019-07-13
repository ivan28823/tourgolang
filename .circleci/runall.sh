#!/bin/bash

SOURCES=$(find -name "*.go" | sort -n)

for o in $SOURCES; do
    echo ">>>>>>>>>>>>>>Running $o"
    go run $o
done

