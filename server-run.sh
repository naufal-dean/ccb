#!/bin/bash
set -euo pipefail

source ./protobuf/protoc.sh

if [[ "$#" -ge 1 ]] && [[ -n "$1" ]]; then
  go run main.go --dbpath ccb.db --port "$1" --name "localhost:$1"
else
  go run main.go --dbpath ccb.db
fi