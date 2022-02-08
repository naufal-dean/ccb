#!/bin/bash

source ./protobuf/protoc.sh

go run main.go --dbpath ccb.db
