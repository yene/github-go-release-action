# GitHub Go Release
Example on how to use Github actions to release a Go binary.

## Checklist
- [ ] build for macos, linux, linux arm
- [ ] build for windows
- [x] include version and Git hash in binary
- [x] compress build with upx

## End result
* An action that when triggered releases for all platforms
* User has to provide version and changelog
* maybe naming like rclone rclone-v1.53.4-osx-amd64.zip
* fuck arm on macos, they can use amd
