coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

migrations_up:
	goose -dir migrations postgres "postgres://postgres:rasik1234@localhost:5432/learning_platform?sslmode=disable" up

migrations_down:
	goose -dir migrations postgres "postgres://postgres:rasik1234@localhost:5432/learning_platform?sslmode=disable" down

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run

creat_migration:
	goose -dir "./migrations" create $(NAME) sql