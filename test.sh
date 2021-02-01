export GIT_COMMIT=$(git rev-parse HEAD)
export INPUT_VERSION=0.1.0
go build -ldflags "-X main.BuildVersion=$INPUT_VERSION -X main.GitCommit=$GIT_COMMIT"
./github-go-release-action version
./github-go-release-action
