#!/usr/bin/env bash

dapr run --app-id client --dapr-grpc-port 50007 --components-path ./components  --config ./config.yaml  -- go run cmd/connect-client/main.go
