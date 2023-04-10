export
DATABASE_URL ?= postgres://postgres:Zydfhm2013@localhost:5432/shop?sslmode=disable
TLS_CERT=./local_cert/cert.pem
TLS_KEY=./local_cert/key.pem

.PHONY: mod
mod:
	go mod tidy
	go mod download

.PHONY: run
run: mod
	CGO_ENABLED=0 go run ./cmd/app

.PHONY: compose-up
compose-up:
	docker compose -f "docker-compose.yaml" up --build -d --force-recreate
	docker compose -f "docker-compose.yaml" logs -f

.PHONY: compose-down
compose-down:
	docker compose -f "docker-compose.yaml" down --remove-orphans

.PHONY: goose-create
goose-create:
	@read -p "Enter migration name:" name; \
	goose -dir "./internal/migrations" create $$name sql

.PHONY: goose-status
goose-status:
	goose -dir "./internal/migrations" postgres "$(DATABASE_URL)" status

.PHONY: goose-up
goose-up:
	goose -dir "./internal/migrations" postgres "$(DATABASE_URL)" up

.PHONY: goose-down
goose-down:
	goose -dir "./internal/migrations" postgres "$(DATABASE_URL)" down

.PHONY: goose-reset
goose-reset:
	goose -dir "./internal/migrations" postgres "$(DATABASE_URL)?sslmode=disable" reset

.PHONY: run-migrate
run-migrate: mod
	CGO_ENABLED=0 go run ./cmd/migrate