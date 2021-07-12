# GitHub Go Release Example Workflow
Example on how a Github action can create a release draft and attach compiled Go binaries.<br>
[.github/workflows/draft-release.yaml](/.github/workflows/draft-release.yaml)

## Features
- [x] Binaries have a version subcommand that prints Version and Git hash.
- [x] build for macOS, linux, linux arm
- [x] build mach-o for macOS amd64+arm64
- [x] build for windows
- [x] include version and Git hash in binary
- [x] compress build with upx
- [x] Reading version from txt file, because it's cool.

## How to make Universal Golang Binary for macOS with Go 1.16
> ! upx does not work with multiple archs

Combine them manually:
```bash
go install github.com/randall77/makefat@latest
GOOS=darwin GOARCH=amd64 go build -o binary-macos-amd64
GOOS=darwin GOARCH=arm64 go build -o binary-macos-arm64
makefat binary-macos binary-macos-amd64 binary-macos-arm64
# or lipo -create -output binary-macos-lipo binary-macos-amd64 binary-macos-arm64
# check
lipo -archs binary-macos
```

## Alternative Snippets
```go
// non flag version
if len(os.Args) >= 2 && os.Args[1] == "version" {
  fmt.Println(strings.TrimSpace(buildVersion), strings.TrimSpace(gitCommit))
  os.Exit(0)
}
```

```bash
# version set with ldflags
GOARM=6 GOOS=linux GOARCH=arm go build -ldflags "-X main.buildVersion=$INPUT_VERSION -X main.gitCommit=$GITHUB_SHA" -o ${NAME}-linux-arm
```

## Other inspirations
* Maybe name the assets like rclone-v1.53.4-osx-amd64.zip
* Create a fat macho binary for macOS with amd+arm
