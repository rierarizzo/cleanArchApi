migration_path := ./database/postgres/migrations

run:
	go run main.go

build:
	go build main.go

sqlc:
	sqlc generate

migration create:
	migrate create -ext sql -dir $(migration_path) $(name)