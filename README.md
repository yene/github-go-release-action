# GitHub Go Release Example Workflow
Example on how a Github action can create a release draft and attach compiled Go binaries.<br>
[.github/workflows/draft-release.yaml](/.github/workflows/draft-release.yaml)

## Features
- [x] Binaries have a version subcommand that prints Version and Git hash.
- [x] build for macOS, linux, linux arm
- [ ] build mach-o for macOS amd64+arm64
- [x] build for windows
- [x] include version and Git hash in binary
- [x] compress build with upx

## How to make Universal Golang Binaries with Go 1.16
Combine them manually:
```
go install github.com/randall77/makefat/
GOOS=darwin GOARCH=amd64 go build -o binary-macos-amd64
GOOS=darwin GOARCH=arm64 go build -o binary-macos-arm64
makefat binary-mac-fat binary-macos-amd64 binary-macos-arm64
```

## Other inspirations
* Maybe name the assets like rclone-v1.53.4-osx-amd64.zip
* Create a fat macho binary for macOS with amd+arm
