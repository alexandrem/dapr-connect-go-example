# Dapr connect-go Example

This repository reproduces the example at https://connect.build/docs/go/getting-started for
Dapr between connect-go and grpc-go implementations.

## Setup

Prerequisites
- Docker
- Dapr
- Go 1.18+

```console
dapr init
# Run the consul instance for discovery
./run-consul.sh
```

We're running consul as a replacement for mDNS which may not work correctly in all environment.

## Running the examples

In a terminal, run either the grpc or connect server with dapr:

```console
./run-grpc-server.sh
```

In a second terminal, run either the grpc or connect client with dapr:
```console
./run-grpc-client.sh
```

The gzip compression issue happens when you use both connect-client and connect-server 
implementations.

## Protobuf Generation

The proto code was generated as such:

### connect-go

```console
buf lint
buf generate
```

### grpc-go

```console
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
greet/v1/greet.proto
```
