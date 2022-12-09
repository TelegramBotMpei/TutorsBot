.PHONY: migrate build up test

build:
	docker-compose build telegram_bot_mpei

up: migrate
	docker-compose up telegram_bot_mpei

migrate:
	migrate -path ./migrations -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@0.0.0.0:5432/${POSTGRES_DB}?sslmode=disable' up

tests:
	go test -v ./...



