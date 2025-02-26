build:
	@go build -o bin/dotx cmd/dotx/main.go

test:
	@go test ./...

cover:
	@go test -cover ./...
