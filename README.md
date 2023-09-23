# Finch

## Set up

1. Install dependencies
- Go
- modd
- goose
- docker and docker compose
- Makefile

2. Copy example env
```
cp .env.example .env
```

3. Add database env variable to source shell file
```
export FINCH_DB_DSN='host=localhost user=finch password=finch dbname=finch sslmode=disable'
```

4. Start program
```sh
make dev
```

5. View application on localhost:4000