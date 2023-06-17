# GitHub Go Release Example Workflow
Example how a Github action can trigger a release on a git commit.
[.github/workflows/release-on-commit.yaml](/.github/workflows/release-on-commit.yaml)

Example on how a Github action can create a release draft, from a manual trigger.<br>
[.github/workflows/draft-release.yaml](/.github/workflows/draft-release.yaml)

## Features
- [x] Binaries have a version subcommand that prints Version and Git hash.
- [x] build for macOS, linux, linux arm
- [x] build mach-o for macOS amd64+arm64
- [x] build for windows
- [x] include version and Git hash in binary
- [x] compress build with upx
- [ ] upx for macos arm (currently upx does not work with mac arm)

## Alternative way to read Git commit from build
After 1.18+ the Git commit is included in the debug information per default and can be read like this:
```Go
	vcsRevision := ""
	info, ok := debug.ReadBuildInfo()
	if ok {
		for _, kv := range info.Settings {
			switch kv.Key {
			case "vcs.revision":
				vcsRevision = kv.Value
			}
		}
	}
	fmt.Println(vcsRevision)
```

## How to make Universal Golang Binary for macOS with Go 1.16
> ! upx does not work with multiple archs

Combine them manually, this obviously doubles the size.
```bash
go install github.com/randall77/makefat@latest
GOOS=darwin GOARCH=amd64 go build -o binary-macos-amd64
GOOS=darwin GOARCH=arm64 go build -o binary-macos-arm64
makefat binary-macos binary-macos-amd64 binary-macos-arm64
# or lipo -create -output binary-macos-lipo binary-macos-amd64 binary-macos-arm64
# check
lipo -archs binary-macos
```

## Other inspirations
* Maybe name the assets like rclone-v1.53.4-osx-amd64.zip
* Create a fat macho binary for macOS with amd+arm

## Links
* https://sanderknape.com/2021/01/go-crazy-github-actions/

## Github action tricks
* stripping V from string https://stackoverflow.com/questions/71630227/github-action-how-to-remove-starting-v-from-variable
*
