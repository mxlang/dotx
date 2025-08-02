MAIN_PACKAGE := "./cmd/dotx"

VERSION := "$(git rev-parse --short HEAD)-dev"
LD_FLAGS := "-s -X main.version=" + VERSION

run *PARAMS: build
	@./bin/dotx {{PARAMS}}

build:
	@go build -ldflags="{{LD_FLAGS}}" -o bin/dotx {{MAIN_PACKAGE}}

install:
	@go install -ldflags="{{LD_FLAGS}}" {{MAIN_PACKAGE}}

test:
	@go test ./...

cover:
	@go test -cover ./...

tidy:
	@go mod tidy

update-deps:
	@go get -u ./...
	@go mod tidy
