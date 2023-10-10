## build
build:
	#clear
	go fmt ./...
	go build -o ./bin/contract_maker ./contract_maker.go

## run
run:
	#clear
	go fmt ./...
	go run contract_maker.go

## lint
lint:
	#clear
	go fmt ./...
	golangci-lint run .
	#gocyclo -over 10 .
	#gocritic check .
	#staticcheck .

## run.test
run.test:
	#clear
	go fmt ./...
	go test -coverprofile cover.out .
	go tool cover -func=cover.out

## mod
mod:
	#clear
	go mod tidy -compat=1.20
	go mod vendor
	go fmt ./...

## help
help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
