build:
	@go build -o bin/dotx ./cmd/dotx

install:
	@go install ./cmd/dotx

test:
	@go test ./...

cover:
	@go test -cover ./...
