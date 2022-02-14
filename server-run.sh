#!/bin/bash

source ./protobuf/protoc.sh

if [ -n "$1" ]; then
  go run main.go --dbpath ccb.db --port "$1"
else
  go run main.go --dbpath ccb.db
fi