#!/usr/bin/env bash

docker run -d -p 8500:8500 --name=dev-consul -e CONSUL_BIND_INTERFACE=eth0 consul
