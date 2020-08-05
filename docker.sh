#!/usr/bin/env bash

set -eo pipefail

cd "$(dirname "$0")"

function die() {
  echo 1>&2 $*
  exit 1
}

if [ -z "$PROTO_SRC" ]; then
  die "must set PROTO_SRC (external proto root directory)"
fi

docker run --rm -it \
  -v "$PROTO_SRC":/proto \
  -v $(pwd):/home \
  -w /home \
  -e PROTO_SRC=/proto \
  asia.gcr.io/sajari9-staging/protoc:v0.0.0-alpha1 \
  ./generate.sh
