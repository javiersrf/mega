#!/usr/bin/env bash

echo "Running tests"
API_PORT=$PORT go test ./...
