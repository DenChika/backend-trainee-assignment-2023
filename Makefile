ifneq (,$(wildcard ./.env))
	include .env
	export
endif

MIGRATIONS_DIR = ./migrations
DATABASE_URL = postgres://postgres:mypassword@localhost:5432/avitoSegmentsDb?sslmode=disable
UP_STEP =
DOWN_STEP = -all
MAIN_GO = cmd/main.go

build:
	docker-compose build

run:
	make build
	docker-compose up

migrate-new:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME)

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database $(DATABASE_URL) up $(UP_STEP)

migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database $(DATABASE_URL) down $(DOWN_STEP)

swag-init:
	swag init -g $(MAIN_GO)