#!/bin/bash

mkdir -p gen/simple

protoc \
  -I pb \
  --go_opt=paths=source_relative \
  --go_out=gen/simple \
  --go-grpc_opt=paths=source_relative \
  --go-grpc_out=gen/simple \
  pb/*.proto