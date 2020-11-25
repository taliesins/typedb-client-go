# client-go
Grakn Go Client

# Pre requisites

## Protoc
Download protoc and put in path: 

[Linux](https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/protoc-3.14.0-linux-x86_64.zip)
[Mac](protoc-3.14.0-osx-x86_64.zip)
[Windows](https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/protoc-3.14.0-win64.zip)

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
protoc --go_out=. protobuf/*.proto
protoc --go_grpc_out=. protobuf/*.proto
```