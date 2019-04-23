#!/bin/sh

# This is just a small sh script to generate the release binaries
export CGO_ENABLED=1
export GOOS=linux
for ARCH in 386 amd64; do
  echo Building iplookup for ${GOOS} + ${ARCH}
  GOARCH=${ARCH} go build -o iplookup-${GOOS}-${ARCH} ..
  sha256sum iplookup-${GOOS}-${ARCH} > iplookup-${GOOS}-${ARCH}.SHA256
  strip iplookup-${GOOS}-${ARCH}
done
