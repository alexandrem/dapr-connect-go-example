#!/usr/bin/env bash

dapr run --app-id server --app-port 8080 --components-path ./components  --config ./config.yaml -- go run cmd/connect-server/main.go
