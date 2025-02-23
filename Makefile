build:
	@go build -o bin/dotx .

test:
	@go test ./...

cover:
	@go test -cover ./...
