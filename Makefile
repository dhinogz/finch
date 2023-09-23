dev:
	docker compose up -d
	@echo "Waiting for db..."
	@sleep 2
	goose -dir="./migrations" postgres "$FINCH_DB_DSN" up
	modd

prod:
	docker compose -f docker-compose.yml -f docker-compose.production.yml up --build

down:
	docker compose down --remove-orphans

down/db:
	docker compose down --remove-orphans -v

migrate:
	goose -dir="./migrations" postgres "$FINCH_DB_DSN" up

