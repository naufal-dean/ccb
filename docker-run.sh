#!/bin/bash
set -euox pipefail

DOCKER_HUB_USER="naufaldean"
IMAGE_NAME="ccb:latest"

if [[ "$#" -ge 1 ]] && [[ -n "$1" ]] && [[ "$1" =~ b.* ]]; then
  docker build --tag "$DOCKER_HUB_USER/$IMAGE_NAME" .
fi

docker run --name ccb-dev -p 50051:50051 -it --rm "$DOCKER_HUB_USER/$IMAGE_NAME" --dbpath ccb.db
