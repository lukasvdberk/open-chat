#!/bin/bash


if $IS_LIVE ; then
    echo 'Running production'
    # TODO implement production probably the same.
else
  go build -o main .
  ./main
fi