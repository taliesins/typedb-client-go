# client-go
A very rough and quick spike to get a Grakn Go Client up and running. Does not follow the official client patterns, but all the functionality seems to work.

# Pre requisites

## Protoc
Download protoc and put in path: 

[Linux](https://github.com/protocolbuffers/protobuf/releases/download/v22.2/protoc-22.2-linux-x86_64.zip)
[Mac](https://github.com/protocolbuffers/protobuf/releases/download/v22.2/protoc-22.2-osx-x86_64.zip)
[Windows](https://github.com/protocolbuffers/protobuf/releases/download/v22.2/protoc-22.2-win64.zip)


For more information see https://developers.google.com/protocol-buffers/docs/downloads

## Protoc Go plugins
Build go protoc plugin
```
export GO111MODULE=on
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
```

Add to path
```
export PATH="$(go env GOPATH)/bin:$PATH"
```

For more information see https://grpc.io/docs/languages/go/quickstart/

# To generate GRPC client

```
protoc --go_out=. v1/protobuf/*.proto
protoc --go-grpc_out=. v1/protobuf/*.proto

protoc --go_out=. v2/protobuf/common/*.proto
protoc --go_out=. v2/protobuf/cluster/*.proto
protoc --go_out=. v2/protobuf/core/*.proto
protoc --go-grpc_out=. v2/protobuf/common/*.proto
protoc --go-grpc_out=. v2/protobuf/cluster/*.proto
protoc --go-grpc_out=. v2/protobuf/core/*.proto
```

# To test

## Pre-requisites
Download TypeDB: 

[Linux](https://github.com/vaticle/typedb/releases/download/2.8.0/typedb-all-linux-2.8.0.tar.gz)
[Mac](https://github.com/vaticle/typedb/releases/download/2.8.0/typedb-all-mac-2.8.0.zip)
[Windows](https://github.com/vaticle/typedb/releases/download/2.8.0/typedb-all-windows-2.8.0.zip)

For more information see https://vaticle.com/download#typedb

## Run TypeDB

Run TypeDB from extracted directory

```
set SERVER_JAVAOPTS="-Xms4G"
set STORAGE_JAVAOPTS="-Xms8G"
SET JAVA_OPTS="-Xms4G -Xmx8G"

CALL typedb.bat server
```

## Run test app

Test up will setup database as required, and will then read and write to it

Compile
```
go build
```

Run
```
typedb-client-go
```