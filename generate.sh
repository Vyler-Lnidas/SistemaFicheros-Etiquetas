#!/bin/zsh
protoc tagFileSystempb/tagFileSystem.proto --go_out=. --go-grpc_out=.
protoc recordpb/recordpb.proto --go_out=. --go-grpc_out=.