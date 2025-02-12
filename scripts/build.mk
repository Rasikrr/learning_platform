coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go fmt ./...
	golangci-lint run

tests:
	go test ./... -v


build:
	mkdir -p ./bin
	go build -o ./bin ./cmd/app/main.go

gen_docs:
	swag init -g internal/ports/http/server.go -d .