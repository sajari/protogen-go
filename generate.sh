#!/bin/bash -e

function die() {
  echo 1>&2 $*
  exit 1
}

if [ -z "$PROTO_SRC" ]; then
  die "must set PROTO_SRC"
fi

echo 1>&2 "Checking required tools:"
for tool in go protoc protoc-gen-go; do
  q=$(which $tool) || die "didn't find $tool"
  echo 1>&2 "$tool: $q"
done

rm -fr sajari/

echo 1>&2 "Building protos:"
for dir in $(find $PROTO_SRC -name '*.proto' | xargs -n1 dirname | sort | uniq); do
  echo 1>&2 "- $dir"
  protoc -I$PROTO_SRC --go_out=paths=source_relative,plugins=grpc:. $dir/*.proto
done
