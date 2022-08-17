#!/usr/bin/env bash

dapr run --app-id server --app-port 50051 --components-path ./components  --config ./config.yaml -- go run cmd/grpc-server/main.go
