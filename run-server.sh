#!/usr/bin/env bash

PORT=$1

if [ -z "$PORT" ]; then
  echo "Usage: ./run-server.sh <PORT>"
  exit 1
fi

echo "Starting MegaSena API on port $PORT..."
go run main.go --port $PORT