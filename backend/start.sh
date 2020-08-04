#!/bin/bash


if $IS_LIVE ; then
    echo 'Running production'
    # TODO implement production probably the same.
else
#  ls **/*.go | entr go build -o main . && ./main
  go build -o main .
  ./main
fi