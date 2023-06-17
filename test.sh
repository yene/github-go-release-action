#!/bin/bash
set -e
export GIT_COMMIT=$(git rev-parse HEAD)
export INPUT_VERSION=0.1.0
go build -ldflags "-X main.buildVersion=$INPUT_VERSION -X main.gitCommit=$GIT_COMMIT"
./github-go-release-action version
./github-go-release-action

GOOS=darwin GOARCH=amd64 go build -o binary-macos-amd64
GOOS=darwin GOARCH=arm64 go build -o binary-macos-arm64
