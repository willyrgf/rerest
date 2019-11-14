#!/usr/bin/env bash

set -x

pGOOS=(freebsd)
GOARCH=amd64

for GOOS in "${pGOOS[@]}"; do
  GOOS="$GOOS" GOARCH="$GOARCH" go build -o bin/rerest-"$GOOS"-"$GOARCH"
done
