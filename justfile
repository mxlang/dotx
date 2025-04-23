run *PARAMS: build
    @./bin/dotx {{PARAMS}}

build:
	@go build -o bin/dotx ./cmd/dotx

test:
	@go test ./...

cover:
	@go test -cover ./...

tidy:
    @go mod tidy

install:
	@go install ./cmd/dotx