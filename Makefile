dev:
	docker compose up -d
	@echo "Waiting for db..."
	@sleep 2
	# migrate -path=./migrations -database=$$FINCH_DB_DSN up
	modd

prod:
	docker compose -f docker-compose.yml -f docker-compose.production.yml up --build

down:
	docker compose down --remove-orphans

down/db:
	docker compose down --remove-orphans -v

migrate:
	migrate -path=./migrations -database=$$FINCH_DB_DSN up

migrate-version:
	migrate -path=./migrations -database=$$FINCH_DB_DSN version

migrate-down:
	migrate -path=./migrations -database=$$FINCH_DB_DSN down

.PHONY: db/migrations/new
migrate-new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}
