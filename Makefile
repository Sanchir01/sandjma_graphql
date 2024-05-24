PHONY:
SILENT:
MIGRATION_NAME ?= new_migration

build:
	go build -o ./.bin/main ./cmd/main/main.go

gql:
	go get github.com/99designs/gqlgen@latest && go run github.com/99designs/gqlgen generate

run: build
	./.bin/main

build-image:
	docker build -t cryptobot-dockerfile .
start-container:
	docker run --name cryptobot-test -p 80:80 --env-file .env cryptobot-dockerfile
migrations-up:
	goose -dir internal/db/migrations postgres "host=localhost user=postgres password=sanchirgarik01 dbname=golangS sslmode=disable" up

migrations-status:
	goose -dir internal/db/migrations postgres "host=localhost user=postgres password=sanchirgarik01 dbname=golangS sslmode=disable" status

migrations-new:
	goose -dir internal/db/migrations create $(MIGRATION_NAME) sql

compose-up:
	docker-compose  up -d
