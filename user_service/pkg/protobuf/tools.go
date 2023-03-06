//go:build tools
// +build tools

package tools

// WARNING: before generation you have to install:
// - https://github.com/protocolbuffers/protobuf/releases
// - go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
// - go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

import (
	// implementation of gRPC in Go
	_ "google.golang.org/grpc"
	// implementation for protocol buffers in Go
	_ "google.golang.org/protobuf/proto"
)
