coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

migrations_up:
	goose -dir migrations postgres "$(DSN)" up

migrations_down:
	goose -dir migrations postgres "$(DSN)" down


lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run