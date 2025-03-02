include .env

create_migration:
	migrate create -ext=sql -dir=internal/migrations -seq $(NAME)

migrate_up:
	migrate -path=./internal/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASS}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up

migrate_down:
	migrate -path=./internal/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASS}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down

run:
	https_proxy=http://127.0.0.1:1087 go run cmd/main.go

.PHONY: create_migration migrate_up migrate_down run