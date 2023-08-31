ifneq (,$(wildcard ./.env))
	include .env
	export
endif

MIGRATIONS_DIR = ./migrations
DATABASE_URL = postgres://postgres:mypassword@localhost:5432/avitoSegmentsDb?sslmode=disable
UP_STEP =
DOWN_STEP = -all

build:
	docker-compose build

run:
	docker-compose up

migrate-new:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME)

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database $(DATABASE_URL) up $(UPSTEP)

migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database $(DATABASE_URL) down $(DOWNSTEP)