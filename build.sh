export GOARCH="amd64"
export GOOS="linux"
export CGO_ENABLED="1"

go build -o doc-service ./cmd/doc_service/main.go