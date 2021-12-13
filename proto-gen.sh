#!/bin/bash

FILES="pb/*.proto"

mkdir -p gen/simple

for file in $FILES; do
  protoc \
    -I pb \
    --go_opt=module=github.com/jhuggart/repro-reflect-failure/gen/simple \
    --go_out=gen/simple \
    --go-grpc_opt paths=source_relative \
    --go-grpc_out=gen/simple \
    $file
done