#export $(shell sed 's/=.*//' .env)
include .env

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

migrations_up:
	goose -dir migrations postgres "$(POSTGRES_DSN)" up

migrations_down:
	goose -dir migrations postgres "$(POSTGRES_DSN)" down

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run

creat_migration:
	goose -dir "./migrations" create $(NAME) sql

print_dsn:
	@echo $(POSTGRES_DSN)
