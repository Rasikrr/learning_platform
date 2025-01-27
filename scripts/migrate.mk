#export $(shell sed 's/=.*//' .env)
-include .env


migrations_up:
	goose -dir migrations postgres "$(POSTGRES_DSN)" up

migrations_down:
	goose -dir migrations postgres "$(POSTGRES_DSN)" down

creat_migration:
	goose -dir "./migrations" create $(NAME) sql

print_dsn:
	@echo $(POSTGRES_DSN)
