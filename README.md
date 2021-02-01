# GitHub Go Release
Example on how to use Github actions to release simple Go binaries.
`binary version` prints version and Git hash.

## Features
- [x] build for macos, linux, linux arm
- [x] build for windows
- [x] include version and Git hash in binary
- [x] compress build with upx

## Other inspirations
* Maybe name the assets like rclone-v1.53.4-osx-amd64.zip
* Create a fat macho binary for macOS with amd+arm
