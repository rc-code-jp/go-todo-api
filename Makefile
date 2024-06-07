init:
	cp .env.template .env
	@make build
	@make up
	docker compose exec app go mod download
	docker compose exec app go mod tidy
	docker compose exec app go build -v ./...
build:
	docker compose build
up:
	docker compose up -d
stop:
	docker compose stop
down:
	docker compose down --remove-orphans
# shell access to container
app:
	docker compose exec app bash
mysql:
	docker compose exec mysql bash
# migrate
migrate-up:
	docker compose exec app migrate -database "mysql://user:password@tcp(mysql:3306)/go-todo" -path /app/src/infrastructure/db/migrations up
migrate-down:
	docker compose exec app migrate -database "mysql://user:password@tcp(mysql:3306)/go-todo" -path /app/src/infrastructure/db/migrations down
