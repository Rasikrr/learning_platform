#export $(shell sed 's/=.*//' .env)
-include .env


migrations_up:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	goose -dir migrations postgres "$(POSTGRES_DSN)" up

migrations_down:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	goose -dir migrations postgres "$(POSTGRES_DSN)" down

migrations_down_to:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	goose -dir migrations postgres "$(POSTGRES_DSN)" down-to $(V)

create_migration:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	goose -dir "./migrations" create $(NAME) sql

print_dsn:
	@echo $(POSTGRES_DSN)
