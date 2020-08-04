#!/bin/bash


if $IS_LIVE ; then
    echo 'Running production'
    # TODO implement production probably the same.
else
    echo "Running development"
    CompileDaemon --build="go build -o main ." --command=./main
fi