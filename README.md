# client-go
A very rough and quick spike to get a Grakn Go Client up and running. Does not follow the official client patterns, but all the functionality seems to work.

# Pre requisites

## Protoc
Download protoc and put in path: 

[Linux](https://github.com/protocolbuffers/protobuf/releases/download/v3.18.1/protoc-3.18.1-linux-x86_64.zip)
[Mac](https://github.com/protocolbuffers/protobuf/releases/download/v3.18.1/protoc-3.18.1-osx-x86_64.zip)
[Windows](https://github.com/protocolbuffers/protobuf/releases/download/v3.18.1/protoc-3.18.1-win64.zip)

For more information see https://developers.google.com/protocol-buffers/docs/downloads

## Protoc Go plugins
Build go protoc plugin
```
export GO111MODULE=on  # Enable module mode
go get google.golang.org/protobuf/cmd/protoc-gen-go 
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

```

Add to path
```
export PATH="$PATH:$(go env GOPATH)/bin"
```

For more information see https://grpc.io/docs/languages/go/quickstart/

# To generate GRPC client

```
protoc --go_out=. v1/protobuf/*.proto
protoc --go_grpc_out=. v1/protobuf/*.proto

protoc --go_out=. v2/protobuf/common/*.proto
protoc --go_out=. v2/protobuf/cluster/*.proto
protoc --go_out=. v2/protobuf/core/*.proto
protoc --go_grpc_out=. v2/protobuf/common/*.proto
protoc --go_grpc_out=. v2/protobuf/cluster/*.proto
protoc --go_grpc_out=. v2/protobuf/core/*.proto
```