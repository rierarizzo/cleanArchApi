run:
	go run main.go

build:
	go build main.go

sqlc:
	sqlc generate

migration create:
	migrate create -ext sql -dir ./database/postgres/migrations $(name)