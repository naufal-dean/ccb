#!/bin/bash
set -euox pipefail

IMAGE_NAME="ccb:multistage"

if [[ "$#" -ge 1 ]] && [[ -n "$1" ]] && [[ "$1" =~ b* ]]; then
  docker build --tag "$IMAGE_NAME" .
fi

docker run --name ccb-dev -p 50051:50051 -it --rm "$IMAGE_NAME" --dbpath ccb.db
