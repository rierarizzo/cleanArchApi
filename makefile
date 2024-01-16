migration_path := database/postgres/migrations

run:
	go run main.go

build:
	go build main.go

sqlc:
	sqlc generate

goose_create:
	goose -dir $(migration_path) create $(name) sql

goose_up:
	goose -dir $(migration_path) up

goose_down:
	goose -dir $(migration_path) down
