#!/bin/bash
set -euo pipefail

CURDIR=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )

protoc --go_out="$CURDIR" --go_opt=paths=source_relative \
    --go-grpc_out="$CURDIR" --go-grpc_opt=paths=source_relative \
    --proto_path="$CURDIR" \
    "$CURDIR/circuitbreaker.proto"

find "$CURDIR" -name \*.pb.go -exec cp {} "$CURDIR/../examples/client/protobuf" \;

# for python
# python3 -m grpc_tools.protoc --proto_path=. --python_out=. --grpc_python_out=. circuitbreaker.proto
