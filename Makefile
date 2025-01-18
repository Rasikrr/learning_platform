coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

migrations_up:
	goose -dir migrations postgres "$(DSN)" up

migrations_down:
	goose -dir migrations postgres "$(DSN)" down
