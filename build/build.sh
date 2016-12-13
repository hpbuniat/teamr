#!/usr/bin/env bash
BASEDIR=$(dirname "$0")
for OS in "freebsd" "linux" "darwin" "windows"; do
  for ARCH in "386" "amd64"; do
    VERSION="$(git describe --tags $1)"
    GOOS=$OS CGO_ENABLED=0 GOARCH=$ARCH go build -ldflags "-X main.Version=$VERSION" -o teamr
    ARCHIVE="teamr-$VERSION-$OS-$ARCH.tar.gz"
    tar -czf $BASEDIR/../dist/$ARCHIVE teamr
    echo $ARCHIVE
  done
done
