#!/bin/bash

test="$( cd "$( dirname "$0" )" && pwd)/../test"

current="$test/current"
new="$test/new"

msg="everything is clean as it is"

if [ ! -d "$current" ]; then
  mkdir "$current"
  if [ ! -d "$new" ]; then
    echo "$msg"
  else
    rm -r "$new"
  fi
else
  if [ ! -d "$new" ] &&  [ -n "$( ls -A "$current" )" ]; then
    rm -rf "${current:?}/"*
  else
    if [ -n "$( ls -A "$current" )" ]; then
       rm -rf "${current:?}/"*
    fi

    rm -r "$new"
  fi
fi