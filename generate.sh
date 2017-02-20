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

repo_prefix="github.com/sajari/go-genproto"

pkg_map=""
pkg_map="$pkg_map,Msajari/rpc/status.proto=$repo_prefix/sajari/rpc"
pkg_map="$pkg_map,Msajari/rpc/empty.proto=$repo_prefix/sajari/rpc"
pkg_map="$pkg_map,Msajari/engine/value.proto=$repo_prefix/sajari/engine"
pkg_map="$pkg_map,Msajari/engine/key.proto=$repo_prefix/sajari/engine"
pkg_map="$pkg_map,Msajari/engine/query/query.proto=$repo_prefix/sajari/engine/query"
pkg_map="$pkg_map,Msajari/engine/query/v1/query.proto=$repo_prefix/sajari/engine/query/v1"
pkg_map="$pkg_map,Msajari/api/query/v1/query.proto=$repo_prefix/sajari/api/query/v1"

echo 1>&2 "Building protos:"
for dir in $(find $PROTO_SRC -name '*.proto' | xargs -n1 dirname | sort | uniq); do
  echo 1>&2 "- $dir"
  protoc -I$PROTO_SRC --go_out=plugins=grpc$pkg_map:. $dir/*.proto
done
