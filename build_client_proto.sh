#!/bin/sh

export PATH="$(go env GOPATH)/bin:$PATH"

protoc \
  --plugin=protoc-gen-go=$(go env GOPATH)/bin/protoc-gen-go \
  --go_out=. \
  --go-grpc_out=. \
  ./proto_client/*.proto