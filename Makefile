.PHONY: migrate build up test

build:
	docker-compose build telegram_bot_mpei

up:
	docker-compose up telegram_bot_mpei

migrate:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up

tests:
	go test -v ./...



